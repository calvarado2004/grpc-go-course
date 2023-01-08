package main

import (
	"context"
	pb "github.com/calvarado2004/grpc-go-course/calculator/proto"
	"log"
)

func (server *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {

	log.Printf("Received Sum RPC")

	return &pb.SumResponse{
		SumResult: in.FirstNumber + in.SecondNumber,
	}, nil
}
