package main

import (
	pb "github.com/calvarado2004/grpc-go-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {

	log.Printf("Avg function was invoked with a streaming request")

	var sum int32
	var count int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&pb.AvgResponse{
				Avg: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		log.Printf("Receiving number: %v\n", req.Number)

		number := req.GetNumber()
		sum += number
		count++

	}
}
