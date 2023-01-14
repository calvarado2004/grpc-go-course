package main

import (
	pb "github.com/calvarado2004/grpc-go-course/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var address = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to %v", err)
	}

	log.Printf("Connected to %v", address)

	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)

	id := createBlog(client)

	readBlog(client, id) // This should work fine
	//readBlog(client, "aNonExistingId") // This will fail

	updateBlog(client, id) // This should work fine

	listBlog(client)

}
