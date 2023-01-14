package main

import (
	"context"
	"fmt"
	pb "github.com/calvarado2004/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	fmt.Println("Updating the blog")

	blog := &pb.Blog{
		Id:       id,
		AuthorId: "Carlos",
		Title:    "My first blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	_, err := c.UpdateBlog(context.Background(), blog)
	if err != nil {
		fmt.Printf("Error happened while updating: %v", err)
	}

	fmt.Println("Blog was updated")

}