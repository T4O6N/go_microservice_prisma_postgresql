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

	msgs, err := channel.Consume("testing", "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("Received a message: %s\n", string(msg.Body))
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}
