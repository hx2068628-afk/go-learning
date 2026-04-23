package main

import (
	"context"
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
	body := "hello world"
	err = ch.PublishWithContext(context.Background(), "", queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	failOnError(err, "Failed to publish a message")
}
