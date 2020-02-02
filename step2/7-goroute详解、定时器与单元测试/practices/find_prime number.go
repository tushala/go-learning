// 查找 1- 8000 的所有 素数
package main

import (
	"fmt"
	"go-learning/step2/1-包、函数、常量、数据类型、字符操作/utils"
)

func putNum(intchan chan int, n int) {
	//将1-8000放入管道
	for i := 1; i <= n; i++ {
		intchan <- i
	}
	close(intchan)
}

func primeNum(intchan chan int, primechan chan int, exitchan chan bool) {
	for {
		num, ok := <-intchan
		if !ok{
			break
		}
		if utils.Is_Prime(num) {
			primechan <- num
		}
	}
	exitchan <- true
}
func main() {
	intchan := make(chan int, 1000)
	primechan := make(chan int, 2000)
	exitchan := make(chan bool, 4) // 4个携程

	go putNum(intchan, 8000)
	// 开启4个携程从管道中取出数并判断是否为素数

	for i := 0; i < 4; i++ {
		go primeNum(intchan, primechan, exitchan)
	}
	go func(){
		for i := 0; i < 4; i++ {
			<-exitchan
		}
		close(primechan)
	}()

	// 当从exitchan取出4个结果即可放心关闭primechan
	for {
		res, ok := <-primechan
		if !ok{
			break
		}
		fmt.Println(res)
	}
}
