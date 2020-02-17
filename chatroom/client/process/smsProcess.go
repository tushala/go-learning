package process

import (
	"encoding/json"
	"fmt"
	"go-learning/chatroom/client/utils"
	"go-learning/chatroom/common/message"
)

type SmsProcess struct {
}

func (self *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	// 创建实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil{
		fmt.Println("json.Marshal err")
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal err")
		return
	}
	// 发送mes
	tf := utils.Transfer{
		Conn:CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("SendGroupMes err ",err)
		return
	}
	return
}
