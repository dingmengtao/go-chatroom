package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

// initPool 初始化redis连接池
func initPool(ipAddress string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲连接数
		MaxActive:   maxActive,   //和redis数据库的最大连接数，0表示没有限制
		IdleTimeout: idleTimeout, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化链接，表示连接哪个redis
			return redis.Dial("tcp", ipAddress)
		},
	}
}
