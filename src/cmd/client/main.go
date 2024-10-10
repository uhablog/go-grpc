package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	hellopb "mygrpc/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "uhablog"
)

var (
	addr = flag.String("addr", "localhost:8000", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	fmt.Println("start gRPC Client.")
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connet: %v", err)
	}
	defer conn.Close()

	c := hellopb.NewGreetingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Hello(ctx, &hellopb.HelloRequest{Name: *name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
