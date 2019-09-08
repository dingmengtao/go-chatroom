package process

import (
	"encoding/json"
	"fmt"
	"net"

	"../../common/message"
	"../model"
	"../utils"
)

//UserProcess 结构体
type UserProcess struct {
	Conn net.Conn
}

//ServerProcessLogin 处理登陆请求的消息
func (userProcess *UserProcess) ServerProcessLogin(msg *message.Message) (err error) {
	//1.从msg中取出data,反序列化为logMsg
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("json.unMarshal err=", err)
	}
	//2.声明返回的信息resMsg
	var resMsg message.Message
	resMsg.Type = message.LoginResMsgType
	//3.声明一个loginResMsg
	var loginResMsg message.LoginResMsg
	// //固定用户id和密码（单用户固定数据）
	// if loginMsg.UserID == 100 && loginMsg.UserPwd == "123456" {
	// 	loginResMsg.Code = 200
	// } else {
	// 	loginResMsg.Code = 500
	// 	loginResMsg.Error = "该用户不存在，请注册后登录。。。"
	// }
	//使用redis数据库内的数据验证
	user, err := model.MyUserDao.CheckLogin(loginMsg.UserID, loginMsg.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOT_EXISTS {
			loginResMsg.Code = 500
			loginResMsg.Error = err.Error()
		} else if err == model.ERROR_USER_PWSSWORD_INCORRECT {
			loginResMsg.Code = 403
			loginResMsg.Error = err.Error()
		} else {
			loginResMsg.Code = 505
			loginResMsg.Error = "服务器内部错误！"
		}
	} else {
		loginResMsg.Code = 200
		fmt.Println(user, "登陆成功")
	}

	//4.序列化loginResMsg
	data, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	resMsg.Data = string(data)
	//5.序列化resMsg
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//发送数据
	tf := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	err = tf.WritePkg([]byte(data))

	return
}

//ServerProcessRegister 处理注册的请求
func (userProcess *UserProcess) ServerProcessRegister(msg *message.Message) (err error) {
	//从msg中取出data反序列化成RegisterMsg
	var registerMsg message.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data), &registerMsg)
	if err != nil {
		fmt.Println("server.process.userProcess.go.ServerProcessRegister.json.Unmarshal() err=", err)
		return
	}
	//声明resMsg，作为返回信息结构
	var resMsg message.Message
	resMsg.Type = message.RegisterResMsgType
	var registerResMsg message.RegisterResMsg
	//将registerMsg的用户信息存入到redis数据库作为注册
	err = model.MyUserDao.Register(&registerMsg.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMsg.Code = 505
			registerResMsg.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMsg.Code = 506
			registerResMsg.Error = "server.process.userProcess.go.ServerProcessRegister.注册时发生未知错误！"
		}
	} else {
		registerResMsg.Code = 200
	}
	//序列化registerResMsg
	data, err := json.Marshal(registerResMsg)
	if err != nil {
		fmt.Println("server.process.userProcess.go.ServerProcessRegister.json.Marshal(registerResMsg) err=", err)
		return
	}
	//将registerResMsg的信息赋值给resMsg.Data
	resMsg.Data = string(data)
	//序列化resMsg
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("server.process.userProcess.go.ServerProcessRegister.json.Marshal(resMsg) err=", err)
		return
	}
	//发送data给客户端
	tf := &utils.Transfer{
		Conn: userProcess.Conn,
	}
	fmt.Println("server.process.userProcess.go.ServerProcessRegister data", data)
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("server.process.userProcess.go.ServerProcessRegister.tf.WritePkg(data) err=", err)
		return
	}
	return
}
