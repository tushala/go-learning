package process3

import "fmt"

var(
	userMgr *UserMgr
)
type UserMgr struct {
	onlineUsers map[int] *UserProcess
}

func init() {
	userMgr = &UserMgr{
		onlineUsers:make(map[int] *UserProcess, 1024),
	}
}

// 在线添加
func (self *UserMgr) AddOnlineUser(up *UserProcess){
	self.onlineUsers[up.UserId] = up
}

// 离线删除
func (self *UserMgr) DelOnlineUser(userId int){
	delete(self.onlineUsers, userId)
}

func (self *UserMgr) GetAllOnlineUser() map[int] *UserProcess{
	return self.onlineUsers
}

func (self *UserMgr) GetOnlineUserByID(userId int) (up *UserProcess, err error){
	up, ok := self.onlineUsers[userId]
	if !ok{
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}