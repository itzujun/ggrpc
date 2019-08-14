package main

import (
	"fmt"
	pb "github.com/itzujun/ggrpc/consul"
	"golang.org/x/net/context"
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

	name := "北京Jack"
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Greeting: name})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}
	fmt.Printf("返回结果: %s", r.Reply)
	fmt.Println("")
	resp, err := c.Add(context.Background(), &pb.IntTrans{Input1: 11, Input2: 22})
	fmt.Println("计算结果:", resp.IntRes)

	loc, err := c.DeviceLocation(context.Background(), &pb.LocationRes{Src: "成都"})
	if err != nil {
		fmt.Println("it is error: ", err.Error())
		return
	}
	fmt.Println("计算结果为: ",loc.Replay)

}
