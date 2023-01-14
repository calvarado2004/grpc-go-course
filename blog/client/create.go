package main

import (
	"context"
	pb "github.com/calvarado2004/grpc-go-course/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("Creating the blog")

	blog := &pb.Blog{
		AuthorId: "Carlos",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	log.Printf("Blog has been created: %s", res.Id)

	return res.Id
}
