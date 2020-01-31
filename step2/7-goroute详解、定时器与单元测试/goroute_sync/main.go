package main

import "fmt"

//import (
//	"fmt"
//	"time"
//)
//
//func calc(intChan, resultChan chan int) {
//	for v := range intChan {
//		flag := true
//		for i := 2; i < v; i++ {
//			if v%i == 0 {
//				flag = false
//				break
//			}
//		}
//		if flag {
//			resultChan <- v
//		}
//	}
//}
//func read(ch chan int) {
//	for {
//		var b int
//		b = <-ch
//		fmt.Println(b)
//		time.Sleep(time.Second)
//	}
//}
//func main() {
//	intChan := make(chan int, 1000)
//	resultChan := make(chan int, 1000)
//
//	for i := 0; i < 100; i++ {
// 		intChan <- i
//	}
//	close(intChan)
//	for i := 0; i < 8; i++ {
//		go calc(intChan, resultChan)
//	}
//	go func(){
//		for v := range resultChan{
//			fmt.Println(v)
//		}
//	}()
//	time.Sleep(time.Second*4)
//
//}
func send(ch chan int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	var a struct{}
	exitChan <- a
}
func recv(ch chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}

	var a struct{}
	exitChan <- a
}

func main() {
	ch := make(chan int, 10)
	exitChan := make(chan struct{}, 10)

	go send(ch, exitChan)
	go recv(ch, exitChan)

	var total int = 0
	for _ = range exitChan{
		total += 1
		fmt.Println(1, total)
		if total == 2{
			close(exitChan)
			break
		}
	}
}
