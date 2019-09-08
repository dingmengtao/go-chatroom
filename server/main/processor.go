package main

import (
	"fmt"
	"io"
	"net"

	"../../common/message"
	"../process"
	"../utils"
)

//Processor 结构体
type Processor struct {
	Conn net.Conn
}

//根据客户端发送的消息类型的不同，进行不同的处理（调用不同的函数）
func (processor *Processor) serverProcessMsg(msg *message.Message) (err error) {
	switch msg.Type {
	case message.LoginMsgType:
		//处理登陆
		up := &process.UserProcess{
			Conn: processor.Conn,
		}
		err = up.ServerProcessLogin(msg)
	case message.RegisterMsgType:
		//处理注册
		up := &process.UserProcess{
			Conn: processor.Conn,
		}
		err = up.ServerProcessRegister(msg)
	default:
		fmt.Println("server.main.processor.go.serverProcessMsg.消息类型未知。。。")
	}
	return
}

func (processor *Processor) processM2() (err error) {
	//读取客户端发送的消息
	for {
		//读取客户端发来的数据，将该功能封装成方法readPkg()
		fmt.Println("读取客户端发送的数据。。。")
		tf := &utils.Transfer{
			Conn: processor.Conn,
		}
		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，我（服务端）随之退出。。。")
				return err
			}
			fmt.Println("readPkg err=", err)
			return err
		}
		//fmt.Println("msg=", msg)
		//return
		err = processor.serverProcessMsg(&msg)
		if err != nil {
			fmt.Println("serverProcessMsg err=", err)
			return err
		}
	}
}
