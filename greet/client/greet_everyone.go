package main

import (
	"context"
	"fmt"
	pb "github.com/calvarado2004/grpc-go-course/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	fmt.Println("Starting to do a GreetEveryone RPC...")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GreetEveryone RPC: %v", err)
	}

	requests := []*pb.GreetRequest{
		{FirstName: "Carlos"},
		{FirstName: "Laura"},
		{FirstName: "Cristina"},
		{FirstName: "Carlitos"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			log.Printf("Sending request %v", req)
			err := stream.Send(req)
			if err != nil {
				log.Fatalf("Error while sending stream: %v", err)
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving stream: %v", err)
				break
			}
			log.Printf("Received: %v", res.GetResult())
		}

		close(waitc)
	}()

	<-waitc
}
