package main

import (

	"os"
	"log"

	. "MQ/common"
	"github.com/streadway/amqp"
)

func main(){
	conn  :=  GetRabbitConn()

	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open an channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_direct",     //name
		"direct",         //type
		true,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to declare an exchange")

	body := BodyFrom(os.Args)
	err = ch.Publish(
		"logs_direct",            // exchange
		SeverityFrom(os.Args),    //routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] sent %s", body)
}

//go run emit_log.go error "this is a log message"

