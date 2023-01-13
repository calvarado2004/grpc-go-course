package main

import (
	pb "github.com/calvarado2004/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var address = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to %v", err)
	}

	log.Printf("Connected to %v", address)

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	//doGreet(client)
	//doGreetManyTimes(client)
	//doLongGreet(client)
	//doGreetEveryone(client)

	doGreetWithDeadline(client, 5*time.Second) // should complete

}
