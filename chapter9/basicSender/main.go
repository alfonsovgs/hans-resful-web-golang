package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "Dialing failed to RabbitMQ broker")

	defer conn.Close()

	channel, err := conn.Channel()
	handleError(err, "Fetching chanel failed")

	defer channel.Close()

	testQueue, err := channel.QueueDeclare(
		"test", // Name of the queue
		false,  // Message is persisted or not
		false,  // Delete mesage when unused
		false,  // Excusive
		false,  // No Waiting time
		nil,    // Extra args
	)

	handleError(err, "Queue creation failed")

	serverTime := time.Now()
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(serverTime.String()),
	}

	// Publish message
	err = channel.Publish(
		"",             //Exchange
		testQueue.Name, // Routing key (Queue)
		false,          // Mandatory
		false,          // Inmediate
		message,
	)
	handleError(err, "Failed to publish a message")
	log.Println("Successfully published a message to the queue")

}
