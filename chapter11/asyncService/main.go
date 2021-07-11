package main

import (
	"context"
	"log"
	"time"

	proto "github.com/alfonsovgs/hands_web_service/chapter11/asyncService/proto"
	"github.com/asim/go-micro/v3"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("weather"),
	)

	pub := micro.NewPublisher("alerts", service.Client())

	go func() {
		for now := range time.Tick(15 * time.Second) {
			log.Println("Publishing weather alert to Topic: alerts")
			pub.Publish(context.TODO(), &proto.Event{
				City:       "Munich",
				Timestamp:  now.UTC().Unix(),
				Temperatre: 2,
			})
		}
	}()

	service.Init()

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
