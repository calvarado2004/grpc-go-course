package main

import (
	"context"
	pb "github.com/calvarado2004/grpc-go-course/calculator/proto"
	"log"
)

func doSum(client pb.CalculatorServiceClient) {
	log.Printf("Starting to do a Sum RPC...")

	res, err := client.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  35,
		SecondNumber: 35,
	})

	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}

	log.Printf("Response from Sum: %v", res.SumResult)

}
