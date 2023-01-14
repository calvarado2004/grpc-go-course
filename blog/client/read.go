package main

import (
	"context"
	"fmt"
	pb "github.com/calvarado2004/grpc-go-course/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	fmt.Println("Reading the blog")

	blogId := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), blogId)
	if err != nil {
		fmt.Printf("Error happened while reading: %v", err)
	}

	fmt.Printf("Blog was read: %v", res)

	return res

}