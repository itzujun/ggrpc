package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
)

const (
	add = "localhost:50051"
)



func main() {

	conn, err := grpc.Dial(add, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)

	name := "北京刘祖军"
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Greeting: name})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}
	fmt.Printf("返回结果: %s", r.Reply)
}

