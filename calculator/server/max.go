package main

import (
	pb "github.com/calvarado2004/grpc-go-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {

	log.Printf("Max function was invoked with a streaming request")

	var maximum int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {

			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		if number := req.Number; number > maximum {
			maximum = number
			err := stream.Send(&pb.MaxResponse{
				Result: maximum,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v", err)
			}
		}
	}
}