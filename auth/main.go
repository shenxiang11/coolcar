package main

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/shenxiang11/coolcar/shared/server"
	"github/shenxiang11/coolcar/auth-service/auth"
	"github/shenxiang11/coolcar/auth-service/dao"
	authpb "github/shenxiang11/coolcar/auth-service/gen/go/proto"
	"github/shenxiang11/coolcar/auth-service/token"
	"github/shenxiang11/coolcar/auth-service/wechat"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var privateKeyFile = "auth/private.key"

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/coolcar"))
	if err != nil {
		logger.Fatal("cannot connect mongodb", zap.Error(err))
	}

	pkFile, err := os.Open(privateKeyFile)
	if err != nil {
		logger.Fatal("cannot open private key", zap.Error(err))
	}

	pkBytes, err := ioutil.ReadAll(pkFile)
	if err != nil {
		logger.Fatal("cannot read private key", zap.Error(err))
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		logger.Fatal("cannot parse private key", zap.Error(err))
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name: "auth",
		Addr: "localhost:10001",
		RegisterFunc: func(server *grpc.Server) {
			authpb.RegisterAuthServiceServer(server, &auth.Service{
				Logger: logger,
				OpenIDResolver: &wechat.Service{
					AppID:     "wxdcaa4b1c9bc60940",
					AppSecret: "c5c6f5ce172a92817f9b5cc757f0436c",
				},
				Mongo:          dao.NewMongo(mongoClient.Database("coolcar")),
				TokenGenerator: token.NewJWTTokenGen("coolcar/auth", privateKey),
				TokenExpire:    30 * time.Minute,
			})
		},
		Logger: logger,
	}))
}
