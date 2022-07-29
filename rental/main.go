package main

import (
	rentalpb "github.com/shenxiang11/coolcar/rental/gen/go/proto"
	"github.com/shenxiang11/coolcar/rental/trip"
	auth "github.com/shenxiang11/coolcar/shared/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	lis, err := net.Listen("tcp", ":10002")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}

	i, err := auth.Interceptor("../shared/auth/public.key")
	if err != nil {
		logger.Fatal("cannot interceptor", zap.Error(err))
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(i))
	rentalpb.RegisterTripServiceServer(s, &trip.Service{
		Logger: logger,
	})

	err = s.Serve(lis)
	if err != nil {
		logger.Fatal("cannot serve", zap.Error(err))
	}
}
