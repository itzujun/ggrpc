package main

import (
	"fmt"
	pb "github.com/itzujun/ggrpc/consul"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"strings"
)

const (
	port = "127.0.0.1:50051"
)


//局域网设备访问需要绑局域网地址
func getIP()string{
	name, err := os.Hostname()
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return ""
	}
	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return ""
	}

	for _, a := range addrs {
		fmt.Println(a)
		if strings.Contains(a,"192.168"){
			return  a
		}
	}
	return ""
}



type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Reply: "你好 " + in.Greeting}, nil
}

func (s *server) Add(ctx context.Context, in *pb.IntTrans) (*pb.ResRes, error) {
	resp := in.Input1 + in.Input2
	return &pb.ResRes{IntRes: resp}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		fmt.Println("监听启动成功")
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
