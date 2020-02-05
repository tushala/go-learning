package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis Dial error=", err)
	}
	fmt.Println("conn succ...", conn)
	defer conn.Close()
	//go操控redis
	//_, err = conn.Do("Set", "name", "tomejerry")
	_, err = conn.Do("HSet", "user1", "name", "tomejerry")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	_, err = conn.Do("HSet", "user1", "age", 18)
	if err != nil {
		fmt.Println("Hset err=", err)
		return
	}

	fmt.Println("操作ok")
	// 读取redis数据

	//r, err := redis.String(conn.Do("Get", "name"))
	//if err != nil{
	//	fmt.Println("get err=", err)
	//	return
	//}
	//fmt.Println(r)

	r1, err := redis.String(conn.Do("HGet", "user1", "name"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}
	fmt.Println(r1)
	r2, err := redis.Int(conn.Do("HGet", "user1", "age"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}
	fmt.Println(r2)
}
