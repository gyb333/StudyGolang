package main

import (
	"github.com/streadway/amqp"
		. "MQ/common"
		"strconv"
)



func main() {
	conn,ch :=GetRabbitConnChan("root","root","Hadoop",5672)
	defer conn.Close()
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"helloQueue", //name
		true,  //durable		//true持久化
		false,  //delete when unused
		false,  //exclusive
		false,  //no wait
		nil,    //arguments
	)
	FailOnError(err, "Failed to declare q queue")

	body := "Hello"
	for i:=0;i<10;i++{
		err = ch.Publish(
			"",     //exchange
			q.Name,     // routing key
			false,  //mandatory
			false, //immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType: "text/plain",
				Body :      []byte(body+strconv.Itoa(i)),
			},
		)
	}


	FailOnError(err, "Failed to publish a message")

}
