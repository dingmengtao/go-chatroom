package message

const (
	//LoginMsgType 登录信息类型
	LoginMsgType = "LoginMsg"
	//LoginResMsgType 登录返回信息类型
	LoginResMsgType = "LoginResMsg"
	//RegisterMsgType 注册信息类型
	RegisterMsgType = "RegisterMsg"
	//RegisterResMsgType 注册返回信息类型
	RegisterResMsgType = "RegisterResMsg"
)

//Message 消息结构体
type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

//LoginMsg 登录消息结构体
type LoginMsg struct {
	UserID    int    `json:"userID"`   //用户名
	UserPwd   string `json:"userPwd"`  //用户密码
	UseraName string `json:"userName"` //用户名
}

//LoginResMsg 登录返回消息结构体
type LoginResMsg struct {
	Code  int    `json:"code"`  //状态码
	Error string `json:"error"` //错误信息
}

//RegisterMsg 注册消息结构体
type RegisterMsg struct {
	//注册
	User User `json:"user"` //类型就是User结构体
}

//RegisterResMsg 注册返回消息结构体
type RegisterResMsg struct {
	//注册响应
	Code  int    `json:"code"`  //返回状态码，400：该用户已存在，200：注册成功
	Error string `json:"error"` //返回错误信息
}
