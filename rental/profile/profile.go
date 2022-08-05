package profile

import (
	"context"
	blobpb "github.com/shenxiang11/coolcar/blob/gen/go/proto"
	rentalpb "github.com/shenxiang11/coolcar/rental-service/gen/go/proto"
	"github.com/shenxiang11/coolcar/rental-service/profile/dao"
	"github.com/shenxiang11/coolcar/shared/auth"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type IdentityResolver interface {
	Resolve(c context.Context, photo []byte) (*rentalpb.Identity, error)
}

type Service struct {
	rentalpb.UnimplementedProfileServiceServer
	BlobClient        blobpb.BlobServiceClient
	PhotoGetExpire    time.Duration
	PhotoUploadExpire time.Duration
	IdentityResolver  IdentityResolver
	Mongo             *dao.Mongo
	Logger            *zap.Logger
}

func (s *Service) GetProfile(c context.Context, req *rentalpb.GetProfileRequest) (*rentalpb.Profile, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	pr, err := s.Mongo.GetProfile(c, aid)
	if err != nil {
		code := s.logAndConvertProfileErr(err)
		if code == codes.NotFound {
			return &rentalpb.Profile{}, nil
		}
		return nil, status.Error(code, "")
	}
	if pr.Profile == nil {
		return &rentalpb.Profile{}, nil
	}
	return pr.Profile, nil
}

func (s *Service) SubmitProfile(c context.Context, req *rentalpb.Identity) (*rentalpb.Profile, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	p := &rentalpb.Profile{
		Identity:       req,
		IdentityStatus: rentalpb.IdentityStatus_PENDING,
	}
	err = s.Mongo.UpdateProfile(c, aid, rentalpb.IdentityStatus_UNSUBMITTED, p)
	if err != nil {
		s.Logger.Error("cannot update profile", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	go func() {
		time.Sleep(15 * time.Second)
		err := s.Mongo.UpdateProfile(context.Background(), aid, rentalpb.IdentityStatus_PENDING, &rentalpb.Profile{
			Identity:       req,
			IdentityStatus: rentalpb.IdentityStatus_VERIFIED,
		})
		if err != nil {
			s.Logger.Error("cannot verify identity", zap.Error(err))
		}
	}()

	return p, nil
}

func (s *Service) ClearProfile(c context.Context, req *rentalpb.ClearProfileRequest) (*rentalpb.Profile, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	p := &rentalpb.Profile{}
	err = s.Mongo.UpdateProfile(c, aid, rentalpb.IdentityStatus_VERIFIED, p)
	if err != nil {
		s.Logger.Error("cannot update profile", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return p, nil
}

func (s *Service) GetProfilePhoto(c context.Context, req *rentalpb.GetProfilePhotoRequest) (*rentalpb.GetProfilePhotoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) CreateProfilePhoto(c context.Context, req *rentalpb.CreateProfilePhotoRequest) (*rentalpb.CreateProfilePhotoResponse, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	br, err := s.BlobClient.CreateBlob(c, &blobpb.CreateBlobRequest{
		AccountId:           aid.String(),
		UploadUrlTimeoutSec: int32(s.PhotoUploadExpire.Seconds()),
	})
	if err != nil {
		s.Logger.Error("cannot update profile photo", zap.Error(err))
		return nil, status.Error(codes.Aborted, "")
	}

	return &rentalpb.CreateProfilePhotoResponse{UploadUrl: br.UploadUrl}, nil
}

func (s *Service) CompleteProfilePhoto(c context.Context, req *rentalpb.CompleteProfilePhotoRequest) (*rentalpb.Identity, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ClearProfilePhoto(c context.Context, req *rentalpb.ClearProfilePhotoRequest) (*rentalpb.ClearProfilePhotoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) logAndConvertProfileErr(err error) codes.Code {
	if err == mongo.ErrNoDocuments {
		return codes.NotFound
	}
	s.Logger.Error("cannot get profile", zap.Error(err))
	return codes.Internal
}
