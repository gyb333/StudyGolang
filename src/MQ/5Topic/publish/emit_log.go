package main

import (
	"github.com/streadway/amqp"
	"os"
	"log"
	. "MQ/common"
)

/*
go run emit_log.go "kern.critical" "A critical kernal error"
go run emit_log.go "kern.test" "A critical kernal.* error"
go run emit_log.go "test.critical" "A critical *.kernal error"
*/

func main() {
	conn := GetRabbitConn()
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	FailOnError(err, "Failed to declare an exchange")

	body := BodyFrom(os.Args)
	err = ch.Publish(
		"logs_topic",          // exchange
		SeverityFrom(os.Args), // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}
