package model

//User 结构体
type User struct {
	UserID   int    `json:"userID"`
	UserPwd  string `json:"userPwd"`
	UserName string `josn:"userName"`
}
