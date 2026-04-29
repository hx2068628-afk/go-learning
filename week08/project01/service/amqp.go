package service

import "github.com/rabbitmq/amqp091-go"

func Amqp() *amqp091.Connection {
	cnn, err := amqp091.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		panic(err)
	}
	return cnn
}
