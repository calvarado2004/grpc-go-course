package main

import (
	"context"
	pb "github.com/calvarado2004/grpc-go-course/calculator/proto"
	"io"
	"log"
)

func doPrimes(c pb.CalculatorServiceClient) {

	log.Printf("doPrimes function was invoked")

	req := &pb.PrimeRequest{
		Number: 12390392840,
	}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Primes RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("Response from Primes: %d", res.Result)
	}

}
