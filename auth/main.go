package main

import (
	"github/shenxiang11/coolcar/auth-service/auth"
	authpb "github/shenxiang11/coolcar/auth-service/gen/go/proto"
	"github/shenxiang11/coolcar/auth-service/wechat"
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

	lis, err := net.Listen("tcp", ":10001")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		Logger: logger,
		OpenIDResolver: &wechat.Service{
			AppID:     "wxdcaa4b1c9bc60940",
			AppSecret: "c5c6f5ce172a92817f9b5cc757f0436c",
		},
	})

	err = s.Serve(lis)
	if err != nil {
		logger.Fatal("cannot serve", zap.Error(err))
	}
}
