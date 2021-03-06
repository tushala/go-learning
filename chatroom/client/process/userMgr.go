package process

import (
	"fmt"
	"go-learning/chatroom/client/model"
	"go-learning/chatroom/common/message"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser // 用户登录成功后对Curuser初始化
func outputOnlineUser() {
	fmt.Println("当前在线用户列表: ")
	for id, user := range onlineUsers {
		//fmt.Println(111, user) todo 有问题
		fmt.Printf("用户id:%d  用户名%s\n", id, user.UserName)
	}
}
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}

	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}
