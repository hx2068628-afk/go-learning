package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") //建立mq连接
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel() //创建通道
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	queue, err := ch.QueueDeclare(
		"hello",
		true,
		false,
		false,
		false,
		nil,
	) //创建队列
	failOnError(err, "Failed to declare a queue")
	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to publish a message")
	// var forever chan struct{}
	func() {
		for d := range msgs {
			fmt.Printf("%s\n", d.Body)
		}
	}()
	// <-forever

}
