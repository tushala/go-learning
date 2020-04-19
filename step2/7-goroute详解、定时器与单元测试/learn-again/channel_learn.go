//https://www.cnblogs.com/wdliu/p/9272220.html
package main

import (
	"fmt"
)

//var readOnlyChan <-chan int  // 只读chan
//var writeOnlyChan chan<- int // 只写chan
//var mychan chan int          //读写chan
//
//mychannel = make(chan int, 10) //定义完成以后需要make来分配内存空间，不然使用会deadlock
//
//read_only := make (<-chan int,10)//定义只读的channel
//write_only := make (chan<- int,10)//定义只写的channel
//read_write := make (chan int,10)//可同时读写

func chan_example1() {
	var ch chan string
	ch <- "wd"    //写
	a, ok := <-ch //读
	if ok {
		fmt.Println(a)
	} else {
		fmt.Println(ok)
	}
	close(ch)
}

func test(c chan int){
	for i:=0;i<10;i++{
		fmt.Println("send ", i)
		c <- i
	}
}
func main() {
	//go chan_example1() 无输出？？
	//time.Sleep(1 * time.Second)

	/*ch :=make(chan string, 1)
	ch <- "wd"    //写
	a, ok := <-ch //读
	if ok {
		fmt.Println(a)
	} else {
		fmt.Println(ok)
	}
	close(ch)*/


	/*mychannel := make(chan int, 10) // 带缓冲区,可以保存多个数据。不带缓冲区channel：只能存一个数据
	for i := 0;i < 10;i++{
		mychannel <- i*2
	}
	close(mychannel)  //关闭管道
	fmt.Println("data lenght: ",len(mychannel))
	for  v := range mychannel {  //循环管道
		fmt.Println(v)
	}
	fmt.Printf("data lenght:  %d",len(mychannel))*/

	ch := make(chan int)
	go test(ch)
	for j := 0; j < 10; j++ {
		fmt.Println("get ", <-ch)
	}

}
