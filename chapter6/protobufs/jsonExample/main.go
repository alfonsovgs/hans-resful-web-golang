package main

import (
	"encoding/json"
	"fmt"

	pb "github.com/alfonsovgs/hands_web_service/chapter6/protobufs/protofiles"
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

	body, _ := json.Marshal(p)
	fmt.Println(string(body))
}
