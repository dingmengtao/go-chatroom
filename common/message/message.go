package message

const (
	LoginMsgType     = "LoginMsg"
	LoginResMsgType  = "LoginResMsg"
	RegisterMsgType  = "RegisterMsg"
	ResterResMsgType = "RegisterResMsg"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

type LoginMsg struct {
	UserID    int    `json:"userID"`   //用户名
	UserPwd   string `json:"userPwd"`  //用户密码
	UseraName string `json:"userName"` //用户名
}

type LoginResMsg struct {
	Code  int    `json:"code"`  //状态码
	Error string `json:"error"` //错误信息
}

type RegisterMsg struct {
	//注册
}

type RegisterResMsg struct {
	//注册响应
}
