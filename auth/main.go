package main

import (
	"context"
	"github/shenxiang11/coolcar/auth-service/auth"
	"github/shenxiang11/coolcar/auth-service/dao"
	authpb "github/shenxiang11/coolcar/auth-service/gen/go/proto"
	"github/shenxiang11/coolcar/auth-service/wechat"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/coolcar"))
	if err != nil {
		logger.Fatal("cannot connect mongodb", zap.Error(err))
	}

	s := grpc.NewServer()
	authpb.RegisterAuthServiceServer(s, &auth.Service{
		Logger: logger,
		OpenIDResolver: &wechat.Service{
			AppID:     "wxdcaa4b1c9bc60940",
			AppSecret: "c5c6f5ce172a92817f9b5cc757f0436c",
		},
		Mongo: dao.NewMongo(mongoClient.Database("coolcar")),
	})

	err = s.Serve(lis)
	if err != nil {
		logger.Fatal("cannot serve", zap.Error(err))
	}
}
