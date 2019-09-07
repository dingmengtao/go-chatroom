package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"../../common/message"
)

//读取数据
func readPkg(conn net.Conn) (msg message.Message, err error) {
	//延时关闭,这里不能关闭，因为调用它的for循环要用，而不是使用一次就关闭了
	//defer conn.Close()

	buf := make([]byte, 8192)
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	fmt.Printf("读取到的buf=%v \n", buf[:4])

	//把buf[:4]转成uint32并计算七=其长度
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	//读取conn中pkgLen长度的消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(buf[:pkgLen] err=", err)
		return
	}
	//把conn总pkgLen长度的消息内容反序列化
	err = json.Unmarshal(buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

//发送数据
func writeMsg(conn net.Conn, data []byte) (err error) {
	//1.发送一个长度给对方
	var dataLen uint32
	dataLen = uint32(len(data))
	var dataLenByte = make([]byte, 4)
	binary.BigEndian.PutUint32(dataLenByte[0:4], dataLen)
	n, err := conn.Write(dataLenByte)
	if err != nil || n != 4 {
		fmt.Println("conn.Write(dataLenByte) err=", err)
		return
	}
	//fmt.Printf("客户端，发送消息的长度是%d 内容是：%v \n", n, string(data))
	//2.发送一个消息给对方
	n, err = conn.Write(data)
	if err != nil || n != int(dataLen) {
		fmt.Println("conn.Write(data) err=", err)
		return
	}
	//休眠20秒
	//time.Sleep(2 * time.Second)
	//fmt.Println("客户端休眠了20s。。。")
	//9.处理服务器端返回的消息
	return
}
