package main
import (
	. "MQ/common"
	"github.com/streadway/amqp"
	"time"
	"log"
)


func main() {
	conn,ch :=GetRabbitConnChan("root","root","Hadoop",5672)
	defer conn.Close()
	defer ch.Close()
	returns := ch.NotifyReturn(make(chan amqp.Return, 1))
	forever := make(chan bool)
	go handleReturn(returns,forever);

	exchangeName := "return_exchange";
	routingKey := "return.qiye";
	// 发送消息
	msg := "Send Msg And Return Listener By RoutingKey : "

	//mandatory 设置不可路由为true,false 路由不可达消息,自动删除
	ch.Publish(exchangeName,routingKey,true, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body :      []byte(msg+routingKey),
		},)


	//mandatory设置不可路由为true,因为消费端路由key不正确,导致触发amqp.Return
	routingErrorKey := "error.qiye";
	ch.Publish(exchangeName,routingErrorKey,true, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body :      []byte(msg+routingErrorKey),
		},)
	<-forever
}

func handleReturn(returns <-chan amqp.Return,forever chan bool)  {
	for{
		ticker := time.NewTicker(10*time.Second)
		select {
		case Return := <-returns:
			log.Println(Return)
			forever<-true

		case <- ticker.C:
			log.Println("time out")
		}
	}
}