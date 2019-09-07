package main

import (
	"fmt"

	"../process"
)

var userID int
var userPwd string

func main() {
	var key int
	var loop = true

	for loop {
		fmt.Println("----------------欢迎登录聊天室----------------")
		fmt.Println("\t\t\t\t\t 1.登录聊天室")
		fmt.Println("\t\t\t\t\t 2.用户注册")
		fmt.Println("\t\t\t\t\t 3.退出聊天室")
		fmt.Println("请输入：1-3")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户id：")
			fmt.Scanf("%d\n", &userID)
			fmt.Println("请输入用户密码：")
			fmt.Scanf("%s\n", &userPwd)
			//进行登录操作，login.go文件中
			up := &process.UserProcess{}
			up.Login(userID, userPwd)
			//loop = false
		case 2:
			fmt.Println("用户注册")
			//loop = false
		case 3:
			fmt.Println("退出聊天室")
			loop = false
		default:
			fmt.Println("你的输入有误，请重新输入！")
		}
	}

}
