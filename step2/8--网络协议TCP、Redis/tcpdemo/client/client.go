package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	reader := bufio.NewReader(os.Stdin) // 终端输入

	// 从终端输入读取每一行发送给服务器
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("readstring err = ", err)
	}

	// 将 line 发送给服务器
	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("conn Write err= ", err)
	}
	fmt.Printf("客户端发送了 %d 字节的数据", n)
}
