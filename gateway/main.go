package main

import (
	"context"
	authpb "gateway/gen/go/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
)

func main() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:  true,
				UseEnumNumbers: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{},
		},
	))

	err := authpb.RegisterAuthServiceHandlerFromEndpoint(
		c, mux, "localhost:8081",
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	)

	if err != nil {
		log.Fatalf("cannot register auth service: %v", err)
	}

	http.ListenAndServe(":8080", mux)
}
