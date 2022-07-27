package auth

import (
	"context"
	"github/shenxiang11/coolcar/auth-service/dao"
	authpb "github/shenxiang11/coolcar/auth-service/gen/go/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	authpb.UnimplementedAuthServiceServer
	Logger         *zap.Logger
	OpenIDResolver OpenIDResolver
	Mongo          *dao.Mongo
}

type OpenIDResolver interface {
	Resolve(code string) (string, error)
}

func (s *Service) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("received", zap.String("code", req.Code))
	openID, err := s.OpenIDResolver.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolve optionid: %v", err)
	}

	accountID, err := s.Mongo.ResolveAccountID(ctx, openID)
	if err != nil {
		s.Logger.Error("cannot resolve account id", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "")
	}

	return &authpb.LoginResponse{
		AccessToken: accountID,
		ExpiresIn:   0,
	}, nil
}
