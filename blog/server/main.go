package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"log"
	"net"

	pb "github.com/calvarado2004/grpc-go-course/blog/proto"
)

var collection *mongo.Collection

var address = "0.0.0.0:50051"

type Server struct {
	pb.BlogServiceServer
}

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Server listening on %v", address)

	server := grpc.NewServer()
	pb.RegisterBlogServiceServer(server, &Server{})
	reflection.Register(server)

	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server over %v", err)
	}

}
