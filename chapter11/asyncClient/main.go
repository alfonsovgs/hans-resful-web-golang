package main

import (
	"context"
	"log"

	"github.com/alfonsovgs/hands_web_service/chapter11/asyncService/proto"
	"github.com/asim/go-micro/v3"
)

func ProcessEvent(ctx context.Context, event *proto.Event) error {
	log.Println("Got alert: ", event)
	return nil
}

func main() {
	// Create a new service
	service := micro.NewService(
		micro.Name("weather_client"),
	)

	// Initialize the client and parse command line flags
	service.Init()

	micro.RegisterSubscriber("alerts", service.Server(), ProcessEvent)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
