package main

import (
	"context"
	pb "github.com/calvarado2004/grpc-go-course/blog/proto"
	"log"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	println("Deleting the blog")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error happened while deleting: %v", err)
	}

	println("Blog was deleted")
}
