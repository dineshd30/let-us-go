package main

import (
	"context"
	"log"
	"net"

	"github.com/dineshd30/golang/grpc-todos/todos"
	"google.golang.org/grpc"
)

type server struct {
	todos.UnimplementedTodosServer
}

func (s *server) Create(ctx context.Context, cr *todos.CreateRequest) (*todos.CreateResponse, error) {
	log.Printf("Creating todo - %+v", cr)
	return &todos.CreateResponse{Error: "", Id: "4321234"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	todos.RegisterTodosServer(s, &server{})
	log.Println("gRPC todos server listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
