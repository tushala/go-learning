package main

import (
	"fmt"
	"go-learning/chatroom/common/message"
	"go-learning/chatroom/server/utils"
	"go-learning/chatroom/server/process"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (self *Processor) serverProcessMes(mes *message.Message) (err error) {
	// 根据客户端发送消息种类的不同，选择调用哪个函数
	switch mes.Type {
	case message.LoginMesType:
		// 创建 UserProcess 实例
		up := &process.UserProcess{
			Conn: self.Conn,
		}
		err = up.ServerProcessLogin(self.Conn, mes)
	case message.RegisterMesType:
	//
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}
func (self *Processor) process2() (err error){
	for {
		// 将读取数据包直接封装成一个函数readPkg()
		tf := &utils.Transfer{
			Conn: self.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出")
			} else {
				fmt.Println("readPkg err ", err)
			}
			return
		}
		err = self.serverProcessMes(&mes)
		if err != nil {
			return
		}
	}
}
