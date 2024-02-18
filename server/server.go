package main

import (
	"context"
	"fmt"
	"grpcimpl/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

// "grpcimpl/proto"
// "github.com/amris/grpcImplementation/proto"// Adjust if needed
type server struct{}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	result := req.A + req.B
	return &proto.AddResponse{A: result}, nil
}
func (s *server) Substract(ctx context.Context, req *proto.SubstractRequest) (*proto.SubstractResponse, error) {
	result := req.A - req.B
	return &proto.SubstractResponse{A: result}, nil
}
func (s *server) Multiple(ctx context.Context, req *proto.MultipleRequest) (*proto.MultipleResponse, error) {
	result := req.A * req.B
	return &proto.MultipleResponse{A: result}, nil
}
func (s *server) Divide(ctx context.Context, req *proto.DivideRequest) (*proto.DivideResponse, error) {
	if req.B == 0 {
		return nil, fmt.Errorf("division by zero")
	}
	quotient := float32(req.A) / float32(req.B)
	return &proto.DivideResponse{A: int32(quotient)}, nil
}

func main() {
	fmt.Println("trying  to create a grpc server")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}
	s := grpc.NewServer()
	proto.RegisterCalculatorSeviceServer(s, &server{})
	log.Println("server is listening on port : 8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}

}
