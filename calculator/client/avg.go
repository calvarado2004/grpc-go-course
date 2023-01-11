package main

import (
	"context"
	pb "github.com/calvarado2004/grpc-go-course/calculator/proto"
	"log"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Printf("Starting Average Client RPC...")

	// we prepare the request
	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Avg: %v", err)
	}

	numbers := []int32{3, 7, 12, 20, 34}

	for _, number := range numbers {
		log.Printf("Sending number: %v", number)
		err := stream.Send(&pb.AvgRequest{
			Number: number,
		})
		if err != nil {
			return
		}

	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v", err)
	}

	log.Printf("The Avg response is: %v", res)
}
