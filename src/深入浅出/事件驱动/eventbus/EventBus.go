package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

/*
使用 channel 取代回调的理由：
	传统的回调方式要求实现某种接口,
   而channel允许在没有接口的情况下在一个简单的函数中注册订阅者
 */


type DataEvent struct {
	Data interface{}
	Topic string
}

type DataChannel chan DataEvent

type DataChannelSlice []DataChannel

type EventBus struct {
	subscribers map[string]DataChannelSlice
	rm sync.RWMutex
}

func (eb *EventBus) Publish(topic string,data interface{}){
	eb.rm.RLock()
	if chans,found :=eb.subscribers[topic];found{
		channels :=append(DataChannelSlice{},chans...)
		go func(data DataEvent,dcs DataChannelSlice){
			for _,ch :=range dcs{
				ch <-data
			}
		}(DataEvent{Data:data,Topic:topic},channels)
	}
	eb.rm.RUnlock()
}

func (eb *EventBus) Subscribe(topic string,ch DataChannel){
	eb.rm.Lock()
	if prev,found :=eb.subscribers[topic];found{
		eb.subscribers[topic]=append(prev,ch)
	}else{
		eb.subscribers[topic]=append([]DataChannel{},ch)
	}
	eb.rm.Unlock()
}


var eb =&EventBus{
	subscribers:map[string]DataChannelSlice{},
}


func printDataEvent(ch string,data DataEvent){
	fmt.Printf("Channel: %s;Topic: %s;DataEvent:%v\n",ch,data.Topic,data.Data)
}

func publisTo(topic string,data string){
	for{
		eb.Publish(topic,data)
		time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
	}
}

func main()  {
	ch1 :=make(chan DataEvent)
	ch2 :=make(chan DataEvent)
	ch3 :=make(chan DataEvent)

	eb.Subscribe("topic1",ch1)
	eb.Subscribe("topic2",ch2)
	eb.Subscribe("topic3",ch3)

	go publisTo("topic1","Hi topic 1")
	go publisTo("topic2","welcome topic 2")
	go publisTo("topic3","welcome topic 3")

	for{
		select{
		case d:=<-ch1:
			go printDataEvent("ch1",d)
		case d:=<-ch2:
			go printDataEvent("ch2",d)
		case d:=<-ch3:
			go printDataEvent("ch3",d)
		}
	}
}
