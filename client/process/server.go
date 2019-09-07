package process

import (
	"fmt"
	"net"
	"os"

	"../utils"
)

//ShowMenu 显示登陆成功后的界面
func ShowMenu() {
	fmt.Println("----------恭喜xxx登陆成功----------")
	fmt.Println("----------1.显示在线用户列表----------")
	fmt.Println("----------2.发送消息----------")
	fmt.Println("----------3.信息列表----------")
	fmt.Println("----------4.退出系统----------")
	fmt.Println("请选择1-4")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确，请重新输入！")
	}
}

//
func serverProcessMsg(conn net.Conn) {
	//创建一个transfer示例不停地读取服务端消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务端发送的消息。。。")
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		//读取到消息，进行下一步处理逻辑
		fmt.Printf("msg=%v\n", msg)
	}
}
