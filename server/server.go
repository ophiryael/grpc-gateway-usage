package main

import (
	"context"
	"log"
	"net"

	gw "github.com/ophiryael/grpc-gateway-usage/gen/go"
	"google.golang.org/grpc"
)

type YourServiceImpl struct {
	gw.UnimplementedYourServiceServer
}

func (y *YourServiceImpl) Echo(ctx context.Context, request *gw.StringMessage) (*gw.StringMessage, error) {
	return &gw.StringMessage{
		Value: "Hello there!@#",
	}, nil
}

func main() {
	// create new gRPC server
	server := grpc.NewServer()
	// register the GreeterServerImpl on the gRPC server
	gw.RegisterYourServiceServer(server, &YourServiceImpl{})
	// start listening on port :8080 for a tcp connection
	if l, err := net.Listen("tcp", ":9090"); err != nil {
		log.Fatal("error in listening on port :9090", err)
	} else {
		// the gRPC server
		if err := server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}
	}
}
