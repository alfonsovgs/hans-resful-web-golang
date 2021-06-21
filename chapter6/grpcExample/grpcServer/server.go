package main

import (
	"context"
	"log"
	"net"

	pb "github.com/alfonsovgs/hands_web_service/chapter6/grpcExample/protofiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// servr is used to create MoneyTransactionServer
type server struct{}

// MakeTransaction implements MoneyTransactionServer.MakeTransaction
func (s *server) MakeTransaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	log.Printf("Got request for money Transfer....")
	log.Printf("Amount: %f, From A/c:%s, To A/c:%s", in.Amount, in.From, in.To)

	// Use in.Amount, in.From, in.To and perfomr transaction logic
	return &pb.TransactionResponse{Confirmation: true}, nil
}

func main() {
	list, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMoneyTransactionServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
