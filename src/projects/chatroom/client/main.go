package main

import (
			"fmt"
	"projects/chatroom/client/process"
)
//定义两个变量，一个表示用户id, 一个表示用户密码
var userId int
var userPwd string
var userName string

func main() {

	for{
		fmt.Println("登陆聊天室")
		fmt.Println("请输入用户的id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%s\n", &userPwd)
		// 完成登录
		//1. 创建一个UserProcess的实例
		up := &process.UserProcess{}
		up.Login(userId, userPwd)
	}


}
