package main

import (
	pb "github.com/calvarado2004/grpc-go-course/calculator/proto"
	"log"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {

	log.Printf("Primes function was invoked with %v", in)

	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			err := stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})
			if err != nil {
				log.Fatalf("Error while sending data to client: %v", err)
				return err
			}
			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}
