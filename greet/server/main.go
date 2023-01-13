package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	opts := []grpc.ServerOption{}

	tls := true

	if tls {
		certFile := "./ssl/server.crt"
		keyFile := "./ssl/server.pem"

		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed to load certificates: %v", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	log.Printf("Server listening on %v", address)

	server := grpc.NewServer(opts...)

	pb.RegisterGreetServiceServer(server, &Server{})

	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server over %v", err)
	}

}
