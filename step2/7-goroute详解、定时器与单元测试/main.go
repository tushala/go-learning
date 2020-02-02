package main

import (
	"fmt"
	"runtime"
)

type student struct {
	name string
}

func practice1() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	fmt.Println(num)
}
func practice2() {
	var stuChan chan *student
	stuChan = make(chan *student, 10)
	stu := student{"tushala"}
	stuChan <- &stu // 写
}
func practice3() {
	var ch chan int
	ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	for {
		b, ok := <-ch
		if ok == false {
			fmt.Println("chan is close")
			break
		}
		fmt.Println(b)
	}
}
func practice4() {
	//var chan1 chan int   //可读可写
	//var chan2 chan<- int //声明只写
	//chan2 = make(chan int, 3)
	//chan2 <- 20
	//num := <-chan2       // 报错
	//var chan3 <-chan int //声明只读

}
func main() {
	//practice1()
	//practice2()
	//practice3()
	practice4()
}
