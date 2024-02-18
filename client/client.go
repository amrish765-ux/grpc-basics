package main

import (
	"context"
	"fmt"
	"grpcimpl/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dialed server:%v", err)
	}
	defer conn.Close()
	client := proto.NewCalculatorSeviceClient(conn)

	addResponse, err := client.Add(context.Background(), &proto.AddRequest{A: 10, B: 20})
	if err != nil {
		log.Fatalf("failed to add number %v", err)
	}
	fmt.Println("add two number %v", addResponse)

	subtract, err := client.Substract(context.Background(), &proto.SubstractRequest{A: 10, B: 20})
	if err != nil {
		log.Fatalf("failed to substract number %v", err)
	}
	fmt.Printf("subtract: %v\n", subtract)

	multiply, err := client.Multiple(context.Background(), &proto.MultipleRequest{A: 10, B: 20})
	if err != nil {
		log.Fatalf("failed to multi number %v", err)
	}
	fmt.Printf("multiply: %v\n", multiply)
	divide, err := client.Divide(context.Background(), &proto.DivideRequest{A: 100, B: 0})
	if err != nil {
		log.Fatalf("failed to divide number %v", err)
	}
	fmt.Printf("divide: %v\n", divide)
}
