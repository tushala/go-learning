//https://www.cnblogs.com/wdliu/p/9272220.html
package main

import (
	"fmt"
	"sync"
	"time"
)

func example1() {
	c := make(chan string)
	go func() {
		c <- " hello"
	}()
	func() {
		word := <-c + " world"
		fmt.Println(word)
	}()
}
func example2() {
	var n sync.WaitGroup
	for i := 0; i < 20; i++ {
		n.Add(1)
		go func(i int, n *sync.WaitGroup) {
			defer n.Done()
			time.Sleep(1 * time.Second)
			fmt.Printf("goroutine %d is running\n", i)
		}(i, &n)
	}
	n.Wait()
}
func cal1(a int, b int, n *sync.WaitGroup) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	defer n.Done() //goroutinue完成后, WaitGroup的计数-1
}

func example3() {
	var go_sync sync.WaitGroup
	for i := 0; i < 10; i++ {
		go_sync.Add(1)
		go cal1(i, i+1, &go_sync)
	}
	go_sync.Wait()
}
func cal2(a int, b int, Exitchan chan bool) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	time.Sleep(time.Second * 2)
	Exitchan <- true
}
func example4() {
	Exitchan := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go cal2(i, i+1, Exitchan)
	}
	for j := 0; j < 10; j++ {
		<-Exitchan
	}
	close(Exitchan)
}

func send(c chan<- int) {
	// 只写
	for i := 0; i < 10; i++ {
		c <- i
	}
}

func get(c <-chan int) {
	// 只读
	for i := range c {
		fmt.Println(i)
	}
}

func send_data(c chan int) {
	for i := 1; i < 10; i++ {
		c <- i
		fmt.Println("send data : ", i)
	}
}
func main() {
	//example1()
	//time.Sleep(1 * time.Second)

	//example2()

	//example3()

	//example4()

	//c := make(chan int)
	//go send(c)
	//go get(c)
	//
	//time.Sleep(time.Second * 1)

	resch := make(chan int, 20)
	strch := make(chan string, 10)

	go send_data(resch)
	strch <- "wd"
	//time.Sleep(time.Second * 1)
	select {
	case a := <-resch:
		fmt.Println("get data : ", a)
	case b := <-strch:
		fmt.Println("get data : ", b)
	default:
		fmt.Println("no channel actvie")
	}
}
