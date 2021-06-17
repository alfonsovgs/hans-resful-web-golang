package main

import (
	"fmt"

	pb "github.com/alfonsovgs/hands_web_service/chapter6/protobufs/protofiles"
	"github.com/golang/protobuf/proto"
)

func main() {

	p := &pb.Person{
		Id:    1234,
		Name:  "Alfonso Vargas",
		Email: "alfonso@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	p1 := &pb.Person{}

	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)

	fmt.Println("Original struct load from protofile:", p)
	fmt.Println("Marshalled proto:", body)
	fmt.Println("Unmarshalled proto:", body)

}
