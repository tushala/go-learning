package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration){
	pool = &redis.Pool{
		MaxIdle:     maxIdle, // 最大空闲链接数
		MaxActive:   maxActive, //和数据库的最大链接数, 0表示没有限制
		IdleTimeout: idleTimeout,// 最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}