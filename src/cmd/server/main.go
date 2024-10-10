package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	hellopb "mygrpc/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
	// 8080ポートのリスナーを作成
	port := 8000
	listner, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		panic(err)
	}

	// gRPCサーバーを作成
	s := grpc.NewServer()

	// gRPCサーバーにGreetingServiceを登録
	hellopb.RegisterGreetingServiceServer(s, NewMyServer())

	// サーバーリフレクションの設定
	reflection.Register(s)

	// 作成したgRPCサーバーを、8080ポートで稼働
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listner)
	}()

	// Ctrl+Cが入力されたらGraceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
