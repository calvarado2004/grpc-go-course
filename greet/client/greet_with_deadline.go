package main

import (
	"context"
	pb "github.com/calvarado2004/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {

	log.Println("Starting to do a GreetWithDeadline RPC...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{

		FirstName: "Laura",
	}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)

		if !ok {
			log.Fatalf("Non gRPC error when calling Greet: %v", err)
		}

		if e.Code() == codes.DeadlineExceeded {
			log.Println("gRPC Timeout was hit! Deadline was exceeded")
			return
		} else {
			log.Fatalf("gRPC error calling Greet: %v", err)
		}

	}

	log.Printf("Response from GreetWithDeadline: %v", res.Result)
}
