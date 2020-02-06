package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	// 循环读取客户端发送的消息
	for {
		buf := make([]byte, 8096)
		fmt.Println("读取客户端输入...")
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn Write fail ", err)
			return
		}
		fmt.Println("读到buf=", buf[:4])

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
		go process(conn)
	}
}
