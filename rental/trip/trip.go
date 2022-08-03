package trip

import (
	"context"
	rentalpb "github.com/shenxiang11/coolcar/rental-service/gen/go/proto"
	"github.com/shenxiang11/coolcar/shared/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	rentalpb.UnimplementedTripServiceServer
	Logger *zap.Logger
}

func (s *Service) CreateTrip(c context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	s.Logger.Info("create trip", zap.String("start", req.Start))
	s.Logger.Info("account id", zap.String("aid", aid))
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) GetTrip(ctx context.Context, in *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetTrips(ctx context.Context, in *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateTrip(ctx context.Context, in *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	//TODO implement me
	panic("implement me")
}
