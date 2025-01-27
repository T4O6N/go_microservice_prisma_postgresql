package main

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	connection, err := amqp091.Dial("amqp://localhost:5672")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	fmt.Println("Successfully connected to RabbitMQ")

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare("testing", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	err = channel.Publish("", "testing", false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Test Message"),
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Queued status: ", queue)

	fmt.Println("Successfully sent message to RabbitMQ")
}

// MQ_HOST
