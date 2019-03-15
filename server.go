package main

import (
	"fmt"
	pb "github.com/itzujun/ggrpc/consul"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = "127.0.0.1:50051"
)

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
