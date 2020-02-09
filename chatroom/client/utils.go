package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"go-learning/chatroom/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	//fmt.Println("读取客户端输入...")
	_, err = conn.Read(buf[:4])
	if err != nil {
		//fmt.Println("conn Write fail ", err)
		err = errors.New("conn Write fail ")
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	// 根据pkgLen 读取消息内容
	//conn.Read(buf[:pkgLen])
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//fmt.Println("conn read fail ", err)
		//err = errors.New("conn read fail ")
		return
	}
	//把pkgLen 反序列化成m message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		//fmt.Println("json.Unmarshal fail ", err)
		return
	}
	//fmt.Println("读到buf=", buf[:4])
	return
}
