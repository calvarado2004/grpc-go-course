package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/calvarado2004/grpc-go-course/greet/proto"
)

var address = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {

	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Server listening on %v", address)

	server := grpc.NewServer()

	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server over %v", err)
	}

}