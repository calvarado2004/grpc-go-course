package main

import (
	"context"
	pb "github.com/calvarado2004/grpc-go-course/greet/proto"
	"log"
)

func (server *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {

	log.Printf("Greet function was invoked with %v", in)

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil

}
