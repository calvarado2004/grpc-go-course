package main

import (
	"context"
	pb "github.com/calvarado2004/grpc-go-course/calculator/proto"
	"io"
	"log"
	"time"
)

func doMax(c pb.CalculatorServiceClient) {

	log.Printf("Starting to client Max RPC...")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Max: %v\n", err)
	}

	waitc := make(chan struct{})

	go func() {

		numbers := []int32{1, 5, 3, 6, 2, 20}

		for _, number := range numbers {
			log.Printf("Sending number: %d\n", number)
			err = stream.Send(&pb.MaxRequest{Number: number})
			if err != nil {
				log.Fatalf("Error while sending number: %v\n", err)
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
				log.Fatalf("Error while receiving data from Max: %v\n", err)
			}
			log.Printf("Received a new maximum: %d\n", res.Result)

		}

		close(waitc)
	}()

	<-waitc
}
