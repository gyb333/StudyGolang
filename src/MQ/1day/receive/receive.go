package main

import (
	"github.com/streadway/amqp"
	"log"
	"fmt"
	. "MQ/common"
)



func main() {
	RabbitUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", "root", "root", "192.168.56.100", 5672)
	conn, err := amqp.Dial(RabbitUrl)
	FailOnError(err, "Failed to connect to server")
	defer conn.Close();

	ch, err := conn.Channel()
	FailOnError(err, "Failed to connect to channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"helloQueue",    //name
		true,      //durable
		false,      //delete when usused
		false,      // exclusive
		false,      //no-wait
		nil,        // arguments
	)

	FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,     // queue
		"",         // consumer
		true,       // auto-ack		true
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // arguments
	)
	FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func(){
		for d:= range msgs{
			log.Printf("Received a message : %s", d.Body)

			//d.Ack(false)		//如果 autoAck 为false 必须手动发送一个确认消息.
		}
	}()

	log.Printf(" [*] Waiting for messages, To exit press CTRL+C")
	<-forever
}


