package main

import (
	"context"
	pb "github.com/calvarado2004/grpc-go-course/greet/proto"
	"log"
)

func doGreet(client pb.GreetServiceClient) {

	log.Printf("Starting to do a Greet RPC...")

	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Carlos",
	})

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)

	}

	log.Printf("Response from Greet: %v", res.Result)

}
