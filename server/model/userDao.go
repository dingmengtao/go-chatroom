package model

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//在服务器启动后创建一个userDao实例，创建成全局变量，在需要和redis操作时，直接使用即可
var (
	MyUserDao *UserDao
)

// UserDao 结构体
type UserDao struct {
	Pool *redis.Pool
}

// NewUserDao 使用工厂模式创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		Pool: pool,
	}
	return
}

// GetUserByID 根据用户id获取用户信息
func (userDao *UserDao) GetUserByID(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOT_EXISTS
		}
		return
	}

	//把res反序列化
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("userDao.GetUserById.json.Unmarshal.err=", err)
		return
	}
	return
}

// CheckLogin 完成登录校验
func (userDao *UserDao) CheckLogin(userID int, userPwd string) (user *User, err error) {
	conn := userDao.Pool.Get()
	defer conn.Close()
	user, err = userDao.GetUserByID(conn, userID)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWSSWORD_INCORRECT
		return
	}
	return
}
