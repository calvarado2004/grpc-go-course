package main

import (
	"context"
	pb "github.com/calvarado2004/grpc-go-course/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("Starting to do a LongGreet RPC client...")

	reqs := []*pb.GreetRequest{
		{FirstName: "Carlos"},
		{FirstName: "Laura"},
		{FirstName: "Cristina"},
		{FirstName: "Carlitos"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req %v\n", req)
		err := stream.Send(req)
		if err != nil {
			return
		}
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}

	log.Printf("LongGreet Response: %v", res)
}
