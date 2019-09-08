package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"../../common/message"
)

//Transfer 这里将佛能根据方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf  []byte
}

//ReadPkg 读取数据
func (transfer *Transfer) ReadPkg() (msg message.Message, err error) {
	//延时关闭,这里不能关闭，因为调用它的for循环要用，而不是使用一次就关闭了
	//defer conn.Close()

	transfer.Buf = make([]byte, 8192)
	_, err = transfer.Conn.Read(transfer.Buf[:4])
	if err != nil {
		fmt.Println("conn.Read err=", err)
		return
	}
	fmt.Printf("读取到的buf=%v \n", transfer.Buf[:4])

	//把buf[:4]转成uint32并计算其长度
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(transfer.Buf[:4])

	//读取conn中pkgLen长度的消息内容
	n, err := transfer.Conn.Read(transfer.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(buf[:pkgLen] err=", err)
		return
	}
	//fmt.Println("server.utils.ReadPkg.transfer.Conn.Read(transfer.Buf[:pkgLen])", n, err, int(pkgLen))
	//把conn总pkgLen长度的消息内容反序列化
	err = json.Unmarshal(transfer.Buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

//WritePkg 发送数据
func (transfer *Transfer) WritePkg(data []byte) (err error) {
	//1.发送一个长度给对方
	var dataLen uint32
	dataLen = uint32(len(data))
	var dataLenByte = make([]byte, 4)
	binary.BigEndian.PutUint32(dataLenByte[0:4], dataLen)
	n, err := transfer.Conn.Write(dataLenByte)
	if err != nil || n != 4 {
		fmt.Println("conn.Write(dataLenByte) err=", err)
		return
	}
	//fmt.Printf("客户端，发送消息的长度是%d 内容是：%v \n", n, string(data))
	//2.发送一个消息给对方
	n, err = transfer.Conn.Write(data)
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
