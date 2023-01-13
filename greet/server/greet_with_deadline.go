package main

import (
	"context"
	pb "github.com/calvarado2004/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline function was invoked with %v", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request!")
		}
		time.Sleep(1 * time.Second)
	}

	firstName := in.FirstName

	result := "Hello " + firstName

	res := &pb.GreetResponse{
		Result: result,
	}

	return res, nil
}
