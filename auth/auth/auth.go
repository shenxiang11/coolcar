package auth

import (
	"context"
	authpb "github/shenxiang11/coolcar/auth-service/gen/go/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	authpb.UnimplementedAuthServiceServer
	Logger         *zap.Logger
	OpenIDResolver OpenIDResolver
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

	return &authpb.LoginResponse{
		AccessToken: openID,
		ExpiresIn:   0,
	}, nil
}
