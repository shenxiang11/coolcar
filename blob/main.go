package main

import (
	"context"
	"github.com/shenxiang11/coolcar/blob/blob"
	"github.com/shenxiang11/coolcar/blob/cos"
	"github.com/shenxiang11/coolcar/blob/dao"
	blobpb "github.com/shenxiang11/coolcar/blob/gen/go/proto"
	"github.com/shenxiang11/coolcar/shared/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		logger.Fatal("cannot connect mongodb", zap.Error(err))
	}
	db := mongoClient.Database("coolcar")

	st, err := cos.NewService("https://coolcar-profile-1253441264.cos.ap-shanghai.myqcloud.com", "AKIDrY0DBuOcbLp43pMN9gxzlIacPwEG2Ztx", "rSovnPmumRr2xd3EArKYWFiwud6KI1Vj")
	if err != nil {
		logger.Fatal("cannot create cos service", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name: "blob",
		Addr: "localhost:10003",
		RegisterFunc: func(server *grpc.Server) {
			blobpb.RegisterBlobServiceServer(server, &blob.Service{
				Storage: st,
				Mongo:   dao.NewMongo(db),
				Logger:  logger,
			})
		},
		Logger: logger,
	}))

}
