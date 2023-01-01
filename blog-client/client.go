package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Ibrahim-Muhammad13/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello form blog client ")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error not connect %v", err)
	}
	defer cc.Close()
	c := blogpb.NewBlogServiceClient(cc)
	blog := &blogpb.Blog{
		AuthorId: "toti",
		Title:    "my first blog",
		Content:  "content of my first blog",
	}
	res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("error calling greet rpc %v", err)
	}
	blogid := res.GetBlog().GetId()
	fmt.Printf("blog created %v \n", res)

	fmt.Println("reading the blog")

	// _, err2 := c.GetBlog(context.Background(), &blogpb.GetBlogRequest{BlogId: "dfsa"})
	// if err2 != nil {
	// 	fmt.Printf("%v", err2)
	// }
	readBlog, errreading := c.GetBlog(context.Background(), &blogpb.GetBlogRequest{BlogId: blogid})
	if errreading != nil {
		fmt.Printf("%v", errreading)
	}
	fmt.Printf("blog was read: %v \n", readBlog)

	//update blog
	newBlog := &blogpb.Blog{
		Id:       blogid,
		AuthorId: "toti updated",
		Title:    "my first blog edited",
		Content:  "content of my first blog edited",
	}
	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if updateErr != nil {
		fmt.Printf("%v", updateErr)
	}
	fmt.Printf("blog was updtaed %v \n", updateRes)

	//delete blog

	deleteBlog, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogid})
	if deleteErr != nil {
		fmt.Printf("%v", deleteErr)
	}
	fmt.Printf("blog was deleted %v", deleteBlog)

	//list blogs

	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		fmt.Printf("%v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("%v", err)
		}
		fmt.Println(res.GetBlog())
	}
}
