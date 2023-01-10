package main

import (
	"context"
	pb "github.com/calvarado2004/grpc-go-course/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {

	log.Println("Starting to do a GreetManyTimes RPC...")

	req := &pb.GreetRequest{
		FirstName: "Carlos",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)

	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			// we've reached the end of the stream
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
	
}