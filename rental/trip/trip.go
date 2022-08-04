package trip

import (
	"context"
	rentalpb "github.com/shenxiang11/coolcar/rental-service/gen/go/proto"
	"github.com/shenxiang11/coolcar/rental-service/trip/dao"
	"github.com/shenxiang11/coolcar/shared/auth"
	"github.com/shenxiang11/coolcar/shared/id"
	"github.com/shenxiang11/coolcar/shared/mongo/objid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
)

type Service struct {
	rentalpb.UnimplementedTripServiceServer
	Mongo          *dao.Mongo
	Logger         *zap.Logger
	ProfileManager ProfileManager
	CarManager     CarManager
	POIManager     POIManager
	DistanceCalc   DistanceCalc
	NowFun         func() int64
}

type ProfileManager interface {
	Verify(ctx context.Context, id id.AccountID) (id.IdentityID, error)
}

type CarManager interface {
	Verify(c context.Context, cid id.CarID, loc *rentalpb.Location) error
	Unlock(c context.Context, cid id.CarID, aid id.AccountID, tid id.TripID, avatarURL string) error
	Lock(c context.Context, cid id.CarID) error
}

type POIManager interface {
	Resolve(context.Context, *rentalpb.Location) (string, error)
}

type DistanceCalc interface {
	DistanceKm(context.Context, *rentalpb.Location, *rentalpb.Location) (float64, error)
}

func (s *Service) CreateTrip(c context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	if req.CarId == "" || req.Start == nil {
		return nil, status.Error(codes.InvalidArgument, "")
	}

	iID, err := s.ProfileManager.Verify(c, aid)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	carID := id.CarID(req.CarId)
	err = s.CarManager.Verify(c, carID, req.Start)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, err.Error())
	}

	ls := s.calcCurrentStatus(c, &rentalpb.LocationStatus{Location: req.Start, TimestampSec: s.NowFun()}, req.Start)
	tr, err := s.Mongo.CreateTrip(c, &rentalpb.Trip{
		AccountId:  aid.String(),
		CarId:      carID.String(),
		Start:      ls,
		Current:    ls,
		Status:     rentalpb.TripStatus_IN_PROGRESS,
		IdentityId: iID.String(),
	})
	if err != nil {
		s.Logger.Warn("cannot create trip", zap.Error(err))
		return nil, status.Error(codes.AlreadyExists, "")
	}

	go func() {
		err := s.CarManager.Unlock(context.Background(), carID, aid, objid.ToTripID(tr.ID), req.AvatarUrl)
		if err != nil {
			s.Logger.Error("cannot unlock car", zap.Error(err))
		}
	}()

	return &rentalpb.TripEntity{
		Id:   tr.ID.Hex(),
		Trip: tr.Trip,
	}, nil
}

func (s *Service) GetTrip(c context.Context, req *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	tr, err := s.Mongo.GetTrip(c, id.TripID(req.Id), aid)
	if err != nil {
		return nil, status.Error(codes.NotFound, "")
	}

	return tr.Trip, nil
}

func (s *Service) GetTrips(c context.Context, req *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	trips, err := s.Mongo.GetTrips(c, aid, req.Status)
	if err != nil {
		s.Logger.Error("cannot get trips", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	res := &rentalpb.GetTripsResponse{}
	for _, tr := range trips {
		res.Trips = append(res.Trips, &rentalpb.TripEntity{
			Id:   tr.ID.Hex(),
			Trip: tr.Trip,
		})
	}

	return res, nil
}

func (s *Service) UpdateTrip(c context.Context, req *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	tid := id.TripID(req.Id)
	tr, err := s.Mongo.GetTrip(c, tid, aid)
	if err != nil {
		return nil, status.Error(codes.NotFound, "")
	}

	if tr.Trip.Status == rentalpb.TripStatus_FINISHED {
		return nil, status.Error(codes.FailedPrecondition, "cannot update a finished trip")
	}

	if tr.Trip.Current == nil {
		s.Logger.Error("trip without current set", zap.String("id", tid.String()))
		return nil, status.Error(codes.Internal, "")
	}

	cur := tr.Trip.Current.Location
	if req.Current != nil {
		cur = req.Current
	}

	tr.Trip.Current = s.calcCurrentStatus(c, tr.Trip.Current, cur)

	if req.EndTrip {
		tr.Trip.End = tr.Trip.Current
		tr.Trip.Status = rentalpb.TripStatus_FINISHED
		err := s.CarManager.Lock(c, id.CarID(tr.Trip.CarId))
		if err != nil {
			return nil, status.Errorf(codes.FailedPrecondition, "cannot lock car: %v", err)
		}
	}

	err = s.Mongo.UpdateTrip(c, tid, aid, tr.UpdateAt, tr.Trip)
	if err != nil {
		return nil, status.Error(codes.Aborted, "")
	}

	return tr.Trip, nil
}

const centsPerSec = 0.7

func (s *Service) calcCurrentStatus(c context.Context, last *rentalpb.LocationStatus, cur *rentalpb.Location) *rentalpb.LocationStatus {
	now := s.NowFun()
	elapsedSec := float64(now - last.TimestampSec)

	dist, err := s.DistanceCalc.DistanceKm(c, last.Location, cur)
	if err != nil {
		s.Logger.Warn("cannot calculate distance", zap.Error(err))
	}

	poi, err := s.POIManager.Resolve(c, cur)
	if err != nil {
		s.Logger.Info("cannot resolve poi", zap.Stringer("loc", cur))
	}

	return &rentalpb.LocationStatus{
		Location:     cur,
		FeeCent:      last.FeeCent + int32(centsPerSec*elapsedSec*2*rand.Float64()),
		KmDriven:     last.KmDriven + dist,
		PoiName:      poi,
		TimestampSec: now,
	}
}
