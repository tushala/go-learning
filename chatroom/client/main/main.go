package main

import (
	"fmt"
	"go-learning/chatroom/client/process"
	"os"
)

var userId int
var userPwd string
var userName string

func main() {
	var key int          // 接受用户选择
	var loop bool = true // 是否还继续显示
	for loop {
		fmt.Println("--------------------欢迎登录多人聊天系统--------------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3): ")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)

			up := &process.UserProcess{}
			up.Login(userId, userPwd)
			//loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户用户名")
			fmt.Scanf("%s\n", &userName)
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
			//loop = false
		case 3:
			fmt.Println("退出系统")
			//loop = false
			os.Exit(0)
		default:
			fmt.Println("输入有误,请重新输入")
		}
	}

	//if key == 1 {
	//	fmt.Println("请输入用户的id")
	//	fmt.Scanf("%d\n", &userId)
	//	fmt.Println("请输入用户的密码")
	//	fmt.Scanf("%s\n", &userPwd)
	//	main2.login(userId, userPwd)
	//	//if err != nil{
	//	//	fmt.Println("登录失败")
	//	//}else{
	//	//	fmt.Println("登录成功")
	//	//}
	//} else {
	//
	//}
}
