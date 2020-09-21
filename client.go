package main

import (
	"log"

	"github.com/sharmarajdaksh/grpc-second-attempt/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := protos.NewChatServiceClient(conn)

	message := protos.Message{
		Body: "Hello world!",
	}

	r, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %v", err)
	}

	log.Printf("Response from Server: %s", r.Body)
}
