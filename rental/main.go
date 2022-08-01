package main

import (
	rentalpb "github.com/shenxiang11/coolcar/rental/gen/go/proto"
	"github.com/shenxiang11/coolcar/rental/trip"
	"github.com/shenxiang11/coolcar/shared/server"
	"google.golang.org/grpc"
	"log"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:              "rental",
		Addr:              "localhost:10002",
		AuthPublicKeyFile: "../shared/auth/public.key",
		RegisterFunc: func(server *grpc.Server) {
			rentalpb.RegisterTripServiceServer(server, &trip.Service{
				Logger: logger,
			})
		},
		Logger: logger,
	}))
}
