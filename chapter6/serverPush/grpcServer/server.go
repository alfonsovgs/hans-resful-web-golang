package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/alfonsovgs/hands_web_service/chapter6/serverPush/protofiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port      = ":50051"
	noOfSteps = 3
)

// server is used to create MoneyTransactionServer
type server struct{}

// MakeTransaction implements MoneyTransactionServer.MakeTransaction
func (s *server) MakeTransaction(in *pb.TransactionRequest, stream pb.MoneyTransaction_MakeTransactionServer) error {
	log.Printf("Go request for money transfer...")
	log.Printf("Amount: $%f, From A/c:%s, To A/c:%s", in.Amount, in.From, in.To)

	// Send streams here
	for i := 0; i < noOfSteps; i++ {
		time.Sleep(time.Second * 2)

		// One task is done, send the successful message
		// back to the client
		r := &pb.TransactionResponse{
			Status:      "good",
			Step:        int32(i),
			Description: fmt.Sprintf("Performing step %d", int32(i)),
		}

		if err := stream.Send(r); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, "status", err)
		}
	}

	log.Printf("Successfully transferred amount $%v from %v to %v", in.Amount, in.From, in.To)
	return nil
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
