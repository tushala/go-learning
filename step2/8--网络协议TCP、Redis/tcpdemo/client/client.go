package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	reader := bufio.NewReader(os.Stdin) // 终端输入
	for {
		// 从终端输入读取每一行发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readstring err = ", err)
		}
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出...")
			break
		}
		// 将 line 发送给服务器
		n, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn Write err= ", err)
		}
		fmt.Printf("客户端发送了 %d 字节的数据\n", n)
	}

}
