package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-learning/chatroom/client/utils"
	"go-learning/chatroom/common/message"
	"net"
	"os"
)

type UserProcess struct {
}

func (self *UserProcess) Login(userId int, userPwd string) (err error) {
	//fmt.Printf("userid= %d userPwd= %s\n", userId, userPwd)
	//return nil
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
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
	if err != nil {
		fmt.Println("json Marshal err= ", err)
	}
	mes.Data = string(data)
	//mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json Marshal err= ", err)
		return
	}
	// 为防止丢包先发送data长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:])
	if n != 4 || err != nil {
		fmt.Println("conn Write fail ", err)
		return
	}
	//fmt.Printf("客户端发送消息长度 = %d 内容是 %s\n", len(data), st ring(data))

	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn Write fail ", err)
		return
	}
	//return
	//time.Sleep(time.Second * 20)
	//fmt.Println("休眠20秒")
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err = ", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		// 启动一个协程保持和服务端的通讯,如果服务器有数据推送给客户端
		go serverProcessMes(conn)
		//fmt.Println("登录成功")
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}
func (self *UserProcess) Register(userId int, userPwd, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net Dial err= ", err)
		return
	}
	defer conn.Close()
	var mes message.Message
	mes.Type = message.RegisterMesType

	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//loginMes 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json Marshal err= ", err)
	}
	mes.Data = string(data)
	//mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json Marshal err= ", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 = ", err)
		return
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("ReadPkg err = ", err)
		return
	}
	var RegisterResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &RegisterResMes)
	if RegisterResMes.Code == 200 {
		fmt.Println("注册成功, 重新登录吧")
	} else {
		fmt.Println(RegisterResMes.Error)
	}
	os.Exit(0)
	return
}
