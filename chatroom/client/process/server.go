package process

import (
	"encoding/json"
	"fmt"
	"go-learning/chatroom/client/utils"
	"go-learning/chatroom/common/message"
	"net"
	"os"
)

func ShowMenu(userId int) {

	fmt.Println("---------恭喜xxx登录成功---------")
	fmt.Println("---------1.显示在线用户列表---------")
	fmt.Println("---------2.发送消息---------")
	fmt.Println("---------3.信息列表---------")
	fmt.Println("---------4.退出系统---------")

	var key int
	var content string
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("发送群息:")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
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
		Conn: Conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息 ")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("ReadPkg err ", err)
			return
		}
		// 如果读取到信息...
		switch mes.Type {
		case message.NotifyUserStatusMesType: // 有人上线了
			var notifyUserStatusMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("json.Unmarshal err")
				return
			}
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType: // 有人群发了
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器收到未知消息类型")
		}
	}
}
