package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-learning/chatroom/common/message"
	"net"
)

type Transfer struct {
	// 分析它应该有哪些字段
	Conn net.Conn
	Buf  [8096]byte // 传输缓冲
}

func (self *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("读取客户端输入...")
	_, err = self.Conn.Read(self.Buf[:4])
	if err != nil {
		//fmt.Println("conn Write fail ", err)
		//err = errors.New("conn Write fail ")
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(self.Buf[0:4])

	// 根据pkgLen 读取消息内容
	//conn.Read(buf[:pkgLen])
	n, err := self.Conn.Read(self.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//fmt.Println("conn read fail ", err)
		//err = errors.New("conn read fail ")
		return
	}
	//把pkgLen 反序列化成m message.Message

	err = json.Unmarshal(self.Buf[:pkgLen], &mes)
	if err != nil {
		//fmt.Println("json.Unmarshal fail ", err)
		return
	}
	//fmt.Println("读到buf=", buf[:4])
	return
}
func (self *Transfer) WritePkg(data []byte) (err error) {
	// 先发送一个长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(self.Buf[0:4], pkgLen)
	n, err := self.Conn.Write(self.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn Write fail ", err)
		return
	}
	// 发送data本身
	n, err = self.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn Write fail ", err)
		return
	}
	return
}
