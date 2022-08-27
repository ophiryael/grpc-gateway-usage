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
	server := grpc.NewServer()

	gw.RegisterYourServiceServer(server, &YourServiceImpl{})

	if l, err := net.Listen("tcp", ":9090"); err != nil {
		log.Fatal("error in listening on port :9090", err)
	} else {
		if err := server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}
	}
}
