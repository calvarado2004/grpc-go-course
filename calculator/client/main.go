package main

import (
	pb "github.com/calvarado2004/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var address = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to %v", err)
	}

	log.Printf("Connected to %v", address)

	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)

	//doSum(client)
	//doPrimes(client)
	//doAvg(client)
	//doMax(client)
	doSqrt(client, 83)

}
