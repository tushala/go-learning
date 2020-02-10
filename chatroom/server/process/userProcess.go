package process3

import (
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
	"go-learning/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	// 字段
	Conn net.Conn
}

func (self *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 注册部分核心代码
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal error = ", err)
		return
	}
	// 成功获取注册信息
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes
	if loginMes.UserId == 100 && loginMes.UserPwd == "12345" {
		// 合法
		loginResMes.Code = 200
		//loginResMes.Error = fmt.Errorf("用户登录成功")
		loginResMes.Error = "用户登录成功"
	} else {
		// 不合法
		loginResMes.Code = 500 // 状态码500
		//loginResMes.Error = fmt.Errorf("该用户不存在，请先注册再使用")
		loginResMes.Error = "该用户不存在，请先注册再使用"
	}
	// 将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.marshal error = ", err)
		return
	}

	resMes.Data = string(data)
	//fmt.Println(67854, resMes)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshal error = ", err)
		return
	}
	// 发送data 封装到writePkg函数
	tf := &utils.Transfer{
		Conn:self.Conn,
	}
	err = tf.WritePkg(data)
	return
}
