package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// 全局pool
var pool *redis.Pool

//初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8, // 最大空闲链接数
		MaxActive:   0, //和数据库的最大链接数, 0表示没有限制
		IdleTimeout: 100,// 最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	// 先从pool中取出一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "tushala")
	if err != nil {
		fmt.Println("Hset err=", err)
		return
	}

	r1, err := redis.String(conn.Do("Get",  "name"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}
	fmt.Println(r1)
}