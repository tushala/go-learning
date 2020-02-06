package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
	"net"
)

func login(userId int, userPwd string) (err error) {
	//fmt.Printf("userid= %d userPwd= %s\n", userId, userPwd)
	//return nil
	conn, err := net.Dial("tcp", "localhost:8889")
	if err !=nil{
		fmt.Println("net Dial err= ", err)
		return
	}
	defer conn.Close()
	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil{
		fmt.Println("json Marshal err= ", err )
	}
	mes.Data = string(data)
	//mes 序列化
	data, err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json Marshal err= ", err )
	}
	// 为防止丢包先发送data长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)
	n, err := conn.Write(buf[:])
	if n!=4 || err !=nil{
		fmt.Println("conn Write fail ", err)
	}
	fmt.Printf("客户端发送消息长度 = %d 内容是 %s\n", len(data), string(data))
	return
}
