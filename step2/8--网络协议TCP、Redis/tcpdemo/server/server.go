package main

import (
	"fmt"
	"net"
)
func process(conn net.Conn){
	// 接受客户端发送的数据
	defer conn.Close()
	for{
		buf := make([]byte, 1024)
		fmt.Printf("服务器等待 %s 输入...\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) // 从conn 读取
		if err != nil{
			//if err == io.EOF{
			//	fmt.Println("客户端已退出")
			//}else{
			//	fmt.Println("服务器 Read err = ", err)
			//}
			return
		}
		// 显示内容至终端
		fmt.Println(string(buf[:n]))
	}
}
func main() {
	fmt.Println("服务器开始监听....")
	// tcp协议, 本地监听, 监听端口8888
	listen, err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil{
		fmt.Printf("listen err = ", err)
		return
	}
	defer listen.Close()
	for{
		// 等待客户端连接
		fmt.Println("等待客户端连接... ")
		conn, err := listen.Accept()
		if err != nil{
			fmt.Printf("Accept err = ", err)
		}else{
			fmt.Printf("Accept suc con= %v\n", conn)
		}
		go process(conn)
	}
}
