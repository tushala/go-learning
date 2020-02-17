package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterType            = "RegisterMes"
	RegisterMesType         = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)
const (
	UserOnline = iota
	UserOffline
	UserBusy
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息的数据
}

// 定义两个消息

type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` // 用户名
}
type LoginResMes struct {
	Code    int    `json:"code"` //状态码 500表示没注册 200表示登录成功
	UsersId []int  //保存用户id的切片
	Error   string `json:"error"` //错误信息
}

type RegisterMes struct {
	User
}
type RegisterResMes struct {
	Code  int    `json:"code"`  //状态码 400表示已经注册 200表示注册成功
	Error string `json:"error"` //错误信息
}

type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

// 发送消息结构体
type SmsMes struct {
	Content string `json:"content"` // 内容
	User
}
