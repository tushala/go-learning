package main

import (
	"fmt"
	"io"
	"net"
)

//func writePkg(conn net.Conn, data []byte) (err error) {
//	// 先发送一个长度
//	var pkgLen uint32
//	pkgLen = uint32(len(data))
//	var buf [4]byte
//	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
//	n, err := conn.Write(buf[:])
//	if n != 4 || err != nil {
//		fmt.Println("conn Write fail ", err)
//		return
//	}
//	// 发送data本身
//	n, err = conn.Write(data)
//	if n != int(pkgLen) || err != nil {
//		fmt.Println("conn Write fail ", err)
//		return
//	}
//	return
//}
//func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
//	// 注册部分核心代码
//	var loginMes message.LoginMes
//	err = json.Unmarshal([]byte(mes.Data), &loginMes)
//	if err != nil {
//		fmt.Println("json.Unmarshal error = ", err)
//		return
//	}
//	// 成功获取注册信息
//	var resMes message.Message
//	resMes.Type = message.LoginResMesType
//
//	var loginResMes message.LoginResMes
//	if loginMes.UserId == 100 && loginMes.UserPwd == "12345" {
//		// 合法
//		loginResMes.Code = 200
//		//loginResMes.Error = fmt.Errorf("用户登录成功")
//		loginResMes.Error = "用户登录成功"
//	} else {
//		// 不合法
//		loginResMes.Code = 500 // 状态码500
//		//loginResMes.Error = fmt.Errorf("该用户不存在，请先注册再使用")
//		loginResMes.Error = "该用户不存在，请先注册再使用"
//	}
//	// 将loginResMes 序列化
//	data, err := json.Marshal(loginResMes)
//	if err != nil {
//		fmt.Println("json.marshal error = ", err)
//		return
//	}
//
//	resMes.Data = string(data)
//	//fmt.Println(67854, resMes)
//	data, err = json.Marshal(resMes)
//	if err != nil {
//		fmt.Println("json.marshal error = ", err)
//		return
//	}
//	// 发送data 封装到writePkg函数
//	err = writePkg(conn, data)
//	return
//}

//func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
//	// 根据客户端发送消息种类的不同，选择调用哪个函数
//	switch mes.Type {
//	case message.LoginMesType:
//		err = serverProcessLogin(conn, mes)
//	case message.RegisterMesType:
//	//
//	default:
//		fmt.Println("消息类型不存在，无法处理...")
//	}
//	return
//}
//func readPkg(conn net.Conn) (mes message.Message, err error) {
//	buf := make([]byte, 8096)
//	fmt.Println("读取客户端输入...")
//	_, err = conn.Read(buf[:4])
//	if err != nil {
//		//fmt.Println("conn Write fail ", err)
//		//err = errors.New("conn Write fail ")
//		return
//	}
//	var pkgLen uint32
//	pkgLen = binary.BigEndian.Uint32(buf[0:4])
//
//	// 根据pkgLen 读取消息内容
//	//conn.Read(buf[:pkgLen])
//	n, err := conn.Read(buf[:pkgLen])
//	if n != int(pkgLen) || err != nil {
//		//fmt.Println("conn read fail ", err)
//		//err = errors.New("conn read fail ")
//		return
//	}
//	//把pkgLen 反序列化成m message.Message
//
//	err = json.Unmarshal(buf[:pkgLen], &mes)
//	if err != nil {
//		//fmt.Println("json.Unmarshal fail ", err)
//		return
//	}
//	//fmt.Println("读到buf=", buf[:4])
//	return
//}
func process2(conn net.Conn) {
	// 处理和客户的通讯
	defer conn.Close()
	// 循环读取客户端发送的消息
	processor := Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		if err == io.EOF {
			fmt.Println("客户端退出")
		}else{
			fmt.Println("客户端和服务器通讯协程错误: ", err)
		}
		return
	}

}
func main() {
	fmt.Println("服务器在8889端口监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net listen err: ", err)
		return
	}
	defer listen.Close()
	// 等待客户端连接
	for {
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept err: ", err)
		}
		go process2(conn)
	}
}
