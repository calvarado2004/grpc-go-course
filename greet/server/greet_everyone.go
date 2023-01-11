package main

import (
	"fmt"
	pb "github.com/calvarado2004/grpc-go-course/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {

	log.Println("GreetEveryone was invoked with a streaming request")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		firstName := req.FirstName

		result := fmt.Sprintf("Hello %v!", firstName)

		err = stream.Send(&pb.GreetResponse{
			Result: result,
		})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v", err)
			return err
		}
	}
}
