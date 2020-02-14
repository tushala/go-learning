package process

import (
	"fmt"
	"go-learning/chatroom/client/utils"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("---------恭喜xxx登录成功---------")
	fmt.Println("---------1.显示在线用户列表---------")
	fmt.Println("---------2.发送消息---------")
	fmt.Println("---------3.信息列表---------")
	fmt.Println("---------4.退出系统---------")

	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("输入有误")
	}
}

func serverProcessMes(Conn net.Conn) {
	// 创建一个transfer实例 不停读取服务器发送的信息
	tf := &utils.Transfer{
		Conn:Conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息 ")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("ReadPkg err ", err)
			return
		}
		// 如果读取到信息...
		fmt.Println(mes)
	}
}