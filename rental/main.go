package main

import (
	"context"
	rentalpb "github.com/shenxiang11/coolcar/rental-service/gen/go/proto"
	"github.com/shenxiang11/coolcar/rental-service/trip"
	"github.com/shenxiang11/coolcar/rental-service/trip/ai"
	"github.com/shenxiang11/coolcar/rental-service/trip/dao"
	"github.com/shenxiang11/coolcar/rental-service/trip/manager/car"
	"github.com/shenxiang11/coolcar/rental-service/trip/manager/poi"
	"github.com/shenxiang11/coolcar/rental-service/trip/manager/profile"
	coolenvpb "github.com/shenxiang11/coolcar/shared/coolenv"
	"github.com/shenxiang11/coolcar/shared/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017"))

	ac, err := grpc.Dial("localhost:18001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal("cannot connect aiservice", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:              "rental",
		Addr:              "localhost:10002",
		AuthPublicKeyFile: "../shared/auth/public.key",
		RegisterFunc: func(server *grpc.Server) {
			rentalpb.RegisterTripServiceServer(server, &trip.Service{
				Mongo:          dao.NewMongo(mongoClient.Database("coolcar")),
				Logger:         logger,
				ProfileManager: &profile.Manager{},
				CarManager:     &car.Manager{},
				POIManager:     &poi.Manager{},
				DistanceCalc: &ai.Client{
					AIClient: coolenvpb.NewAIServiceClient(ac),
				},
				NowFun: func() int64 {
					return time.Now().Unix()
				},
			})
		},
		Logger: logger,
	}))
}
