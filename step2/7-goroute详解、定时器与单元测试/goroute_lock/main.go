// 加锁方法很土鳖，引出channel
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m    = make(map[int]uint64) // 全局变量共享资源多个线程操作需要加锁
	lock sync.Mutex
)

type task struct {
	n int
}

func calc(self *task) {
	var mul uint64 = 1
	for i := 1; i < self.n; i++ {
		mul *= uint64(i)
	}
	lock.Lock()
	m[self.n] = mul
	lock.Unlock()
}
func main() {
	for i := 0; i < 1000; i++ {
		t := &task{n: i}
		go calc(t)
	}
	time.Sleep(time.Second * 10)
	for k, v := range m {
		fmt.Println(k, "\t", v)
	}
}
