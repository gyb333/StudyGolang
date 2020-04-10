package common

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"fmt"
)

var Pool *redis.Pool

//建立连接池
func Init(network, address string)  {
	Pool = &redis.Pool{
		MaxIdle:     8,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(network, address)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
}


func FailOnError(err error, msg string){
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

func GetRedisConnection(network, address string)  redis.Conn{

	conn,err:=redis.Dial(network,address)
	FailOnError(err,"Failed to connect to Redis!")
	return conn
} 
