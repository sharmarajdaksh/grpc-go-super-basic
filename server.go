package main

import (
	"log"
	"net"

	"github.com/sharmarajdaksh/grpc-second-attempt/protos"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := protos.Server{}
	gs := grpc.NewServer()

	protos.RegisterChatServiceServer(gs, &s)

	if err := gs.Serve(lis); err != nil {
		log.Fatalf("Failed to server gRPC server over port 9000: %v", err)
	}
}
