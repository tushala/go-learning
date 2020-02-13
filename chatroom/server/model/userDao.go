package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"go-learning/chatroom/common/message"
)
var(
	MyUserDao *UserDao
)
type UserDao struct {
	 pool *redis.Pool
}
func NewUserDao(pool *redis.Pool) (userDao *UserDao){
	userDao = &UserDao{
		pool:pool,
	}
	return
}
func (self * UserDao) getUserById(conn redis.Conn, id int) (user * User, err error) {
	res, err := redis.String(conn.Do("Hget","users", id))
	if err != nil{
		if err == redis.ErrNil{
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil{
		fmt.Println("json Unmarshal err = ", err)
		return
	}
	return
}
func (self *UserDao) Login(userId int, userPwd string)(user *User, err error){
	conn := self.pool.Get()
	defer conn.Close()
	user, err = self.getUserById(conn, userId)
	if err != nil{
		return
	}
	if user.UserPwd != userPwd{
		err = ERROR_USER_PWD
		return
	}
	return
}
func (self *UserDao) Register(user *message.User)(err error){
	conn := self.pool.Get()
	defer conn.Close()
	_, err = self.getUserById(conn, user.UserId)
	if err != nil{
		err = ERROR_USER_EXISTS
		return
	}
	data, err := json.Marshal(user)
	if err != nil{
		fmt.Println("json.Marshal err ", err)
		return
	}
	// 入库
	_, err = conn.Do("HSET", "users", user.UserId, string(data))
	if err != nil{
		fmt.Println("注册用户错误 err = ", err)
		return
	}
	return
}