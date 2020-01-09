package main

import (
			"rpc/rpcdemo"
	"fmt"
	"utils"
)

func main() {
	client,err :=utils.NewClient(":5566")
	if err!=nil{
		panic(err)
	}
	var result float64
	err =client.Call("DemoService.Div",
		rpcdemo.Args{X:3,Y:5},&result)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}
}
