package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"../../common/message"
	"../utils"
)

//UserProcess 结构体
type UserProcess struct {
}

//Login 登录
func (userProcess *UserProcess) Login(userID int, userPwd string) (err error) {
	// fmt.Printf("你的用户ID是：%d，你的用户密码是：%s\n", userID, userPwd)
	// return nil
	//1.连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dail err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var msg message.Message
	msg.Type = message.LoginMsgType
	//3.创建一个LoginMsg结构体
	var loginMsg message.LoginMsg
	loginMsg.UserID = userID
	loginMsg.UserPwd = userPwd
	//4.将loginMsg进行序列化
	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//5.将data赋值给msg.Data
	msg.Data = string(data)
	//6.将msg序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json Marshal err=", err)
	}
	//7.此时的data就是要发送的json消息
	//7.1先把data的长度发送给服务器
	//先获取data长度，再将长度转为byte切片
	var dataLen uint32
	dataLen = uint32(len(data))
	var dataLenByte = make([]byte, 4)
	binary.BigEndian.PutUint32(dataLenByte[0:4], dataLen)
	//7.2再把长度切片发送给服务器
	n, err := conn.Write(dataLenByte)
	if err != nil || n != 4 {
		fmt.Println("conn.Write(dataLenByte) err=", err)
		return
	}
	fmt.Printf("客户端，发送消息的长度是%d 内容是：%v \n", n, string(data))

	//8.发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err=", err)
		return
	}
	//休眠20秒
	//time.Sleep(2 * time.Second)
	//fmt.Println("客户端休眠了20s。。。")
	//9.处理服务器端返回的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	msg, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err=", err)
	}
	var loginResMsg message.LoginResMsg
	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)
	if err != nil {
		fmt.Println("json.Unmarshal() err=", err)
	} else if loginResMsg.Code == 200 {
		//fmt.Println("登陆成功！")
		//开启一个客户端协程，保持和服务器的通讯，如果服务端有数据推送给客户端，则接受并显示在客户端的终端
		go serverProcessMsg(conn)
		//显示登陆后的菜单
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMsg.Error)
	}
	fmt.Println("loginResMsg=", loginResMsg)
	return
}
