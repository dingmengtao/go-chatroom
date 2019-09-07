package main

import (
	"fmt"
	"net"
	"time"

	"../model"
)

// //读取数据
// func readPkg(conn net.Conn) (msg message.Message, err error) {
// 	//延时关闭,这里不能关闭，因为调用它的for循环要用，而不是使用一次就关闭了
// 	//defer conn.Close()

// 	buf := make([]byte, 8192)
// 	_, err = conn.Read(buf[:4])
// 	if err != nil {
// 		fmt.Println("conn.Read err=", err)
// 		return
// 	}
// 	fmt.Printf("读取到的buf=%v \n", buf[:4])

// 	//把buf[:4]转成uint32并计算七=其长度
// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buf[:4])

// 	//读取conn中pkgLen长度的消息内容
// 	n, err := conn.Read(buf[:pkgLen])
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Read(buf[:pkgLen] err=", err)
// 		return
// 	}
// 	//把conn总pkgLen长度的消息内容反序列化
// 	err = json.Unmarshal(buf[:pkgLen], &msg)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal err=", err)
// 		return
// 	}
// 	return
// }

// //发送数据
// func writePkg(conn net.Conn, data []byte) (err error) {
// 	//1.发送一个长度给对方
// 	var dataLen uint32
// 	dataLen = uint32(len(data))
// 	var dataLenByte = make([]byte, 4)
// 	binary.BigEndian.PutUint32(dataLenByte[0:4], dataLen)
// 	n, err := conn.Write(dataLenByte)
// 	if err != nil || n != 4 {
// 		fmt.Println("conn.Write(dataLenByte) err=", err)
// 		return
// 	}
// 	//fmt.Printf("客户端，发送消息的长度是%d 内容是：%v \n", n, string(data))
// 	//2.发送一个消息给对方
// 	n, err = conn.Write(data)
// 	if err != nil || n != int(dataLen) {
// 		fmt.Println("conn.Write(data) err=", err)
// 		return
// 	}
// 	//休眠20秒
// 	//time.Sleep(2 * time.Second)
// 	//fmt.Println("客户端休眠了20s。。。")
// 	//9.处理服务器端返回的消息
// 	return
// }

// //处理登陆请求的消息
// func serverProcessLogin(conn net.Conn, msg *message.Message) (err error) {
// 	//1.从msg中取出data,反序列化为logMsg
// 	var loginMsg message.LoginMsg
// 	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
// 	if err != nil {
// 		fmt.Println("json.unMarshal err=", err)
// 	}
// 	//2.声明返回的信息resMsg
// 	var resMsg message.Message
// 	resMsg.Type = message.LoginResMsgType
// 	//3.声明一个loginResMsg
// 	var loginResMsg message.LoginResMsg
// 	//固定用户id和密码（单用户固定数据）
// 	if loginMsg.UserID == 100 && loginMsg.UserPwd == "123456" {
// 		loginResMsg.Code = 200
// 	} else {
// 		loginResMsg.Code = 500
// 		loginResMsg.Error = "该用户不存在，请注册后登录。。。"
// 	}
// 	//4.序列化loginResMsg
// 	data, err := json.Marshal(loginResMsg)
// 	if err != nil {
// 		fmt.Println("json.Marshal err=", err)
// 		return
// 	}
// 	resMsg.Data = string(data)
// 	//5.序列化resMsg
// 	data, err = json.Marshal(resMsg)
// 	if err != nil {
// 		fmt.Println("json.Marshal err=", err)
// 		return
// 	}
// 	//发送数据
// 	err = writePkg(conn, []byte(data))

// 	return
// }

// //根据客户端发送的消息类型的不同，进行不同的处理（调用不同的函数）
// func serverProcessMsg(conn net.Conn, msg *message.Message) (err error) {
// 	switch msg.Type {
// 	case message.LoginMsgType:
// 		//处理登陆
// 		err = serverProcessLogin(conn, msg)
// 	case message.RegisterMsgType:
// 		//处理注册
// 	default:
// 		fmt.Println("消息类型未知。。。")
// 	}
// 	return
// }

func processM1(conn net.Conn) {
	//延时关闭
	defer conn.Close()
	//读取客户端发送的消息
	p := &Processor{
		Conn: conn,
	}
	err := p.processM2()
	if err != nil {
		fmt.Println("客户端与服务端协程出错 err=", err)
	}
}

//init 初始化
func initAll() {
	//初始化连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	//初始化UserDao
	initUserDao()
}

//初始化UserDao
func initUserDao() {
	//参数pool是redis.go文件内定义的全局变量，因此需要在main方法里先初始化连接池，再初始化UserDao
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	//初始化
	initAll()

	fmt.Println("服务器再监听8889端口")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Lister err=", err)
		return
	}
	//一旦建立监听成功，就等待客户端来连接
	for {
		fmt.Println("等待客户端来连接服务器。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("lister.Accept err=", err)
		}

		//一旦连接成功，则启动一个协程和客户端通信
		go processM1(conn)
	}
}
