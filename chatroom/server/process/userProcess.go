package process3

import (
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
	"go-learning/chatroom/server/model"
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
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			loginResMes.Code = 500
		} else if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 403
		} else {
			loginResMes.Code = 505
		}
		loginResMes.Error = fmt.Sprintf("%s", err)
		fmt.Printf("%s\n", err)
	} else {
		loginResMes.Code = 200
		loginResMes.Error = fmt.Sprintf("%s 登录成功\n", user.UserName)
		fmt.Printf("%s 登录成功\n", user.UserName)
	}

	//if loginMes.UserId == 100 && loginMes.UserPwd == "12345" {
	//	// 合法
	//	loginResMes.Code = 200
	//	//loginResMes.Error = fmt.Errorf("用户登录成功")
	//	loginResMes.Error = "用户登录成功"
	//} else {
	//	// 不合法
	//	loginResMes.Code = 500 // 状态码500
	//	//loginResMes.Error = fmt.Errorf("该用户不存在，请先注册再使用")
	//	loginResMes.Error = "该用户不存在，请先注册再使用"
	//}
	// 将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.marshal error = ", err)
		return
	}

	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshal error = ", err)
		return
	}
	// 发送data 封装到writePkg函数
	tf := &utils.Transfer{
		Conn: self.Conn,
	}
	err = tf.WritePkg(data)
	return
}

func (self *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal error = ", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterMesType
	var registerResMes message.RegisterResMes
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	}else{
		registerResMes.Code = 200
	}
	registerResMes.Error = fmt.Sprintf("%s", err)
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.marshal error = ", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.marshal error = ", err)
		return
	}
	// 发送data 封装到writePkg函数
	tf := &utils.Transfer{
		Conn: self.Conn,
	}
	err = tf.WritePkg(data)
	return
}
