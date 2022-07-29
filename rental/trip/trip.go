package trip

import (
	"context"
	rentalpb "github.com/shenxiang11/coolcar/rental/gen/go/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	rentalpb.UnimplementedTripServiceServer
	Logger *zap.Logger
}

func (s *Service) CreateTrip(c context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error) {
	s.Logger.Info("create trip", zap.String("start", req.Start))
	return nil, status.Error(codes.Unimplemented, "")
}
