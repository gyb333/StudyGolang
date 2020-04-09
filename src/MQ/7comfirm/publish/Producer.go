package main

import (
			. "MQ/common"
	"github.com/streadway/amqp"
		"log"
	"time"
)



func main() {
	conn,ch :=GetRabbitConnChan("root","root","Hadoop",5672)
	defer conn.Close()
	defer ch.Close()

	if err := ch.Confirm(false); err != nil {
		FailOnError(err,"Channel could not be put into confirm mode: %s")
	}
	confirms := ch.NotifyPublish(make(chan amqp.Confirmation, 1))
	forever := make(chan bool)
	go confirm(confirms,forever)


	exchangeName := "confirm_exchange";
	routingKey := "confirm.qiye";

	ch.Publish(exchangeName,routingKey,false,false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body :      []byte("Send Msg By Confirm ..."),
		},)

	<-forever
}



func confirm(confirms <-chan amqp.Confirmation,forever chan bool){
	for{
		ticker := time.NewTicker(10*time.Millisecond)
		select {
		case confirm := <-confirms:
			if confirm.Ack {
				log.Println("confirmed delivery with delivery tag: %d", confirm.DeliveryTag)
				forever<-true
			}
		case <- ticker.C:
			log.Println("time out")
		}
	}


}
