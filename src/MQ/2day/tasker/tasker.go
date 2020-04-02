package main

import (
	"github.com/streadway/amqp"
	"os"
	"log"
	. "MQ/common"
)

//go run tasker.go  First message.
//go run tasker.go  Second message..
//go run tasker.go  Third message...

func main() {
	conn:=GetRabbitConn()
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()


	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	body := BodyFrom(os.Args)
	err = ch.Publish(
		"",           // exchange
		q.Name,       // routing key
		false,        // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
