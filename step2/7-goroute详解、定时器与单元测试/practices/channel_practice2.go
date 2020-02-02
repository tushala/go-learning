package main

import "fmt"

func putNum(intchan chan int, n int) {
	//将1-8000放入管道
	for i := 1; i <= n; i++ {
		intchan <- i
	}
	close(intchan)
}
func calc(intChan chan int, resChan chan map[int]int, exitChan chan bool) {
	for {
		num, ok := <-intChan
		if !ok {
			break
		} else {
			a := make(map[int]int, 1)
			a[num] = int(num * (num + 1) / 2)
			resChan <- a
		}
	}
	exitChan <- true
}
func main() {
	intChan := make(chan int, 200)
	resChan := make(chan map[int]int, 200)
	exitChan := make(chan bool, 4)
	go putNum(intChan, 2000)
	for i := 0; i < 4; i++ {
		go calc(intChan, resChan, exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(resChan)
	}()
	for {
		res, ok := <-resChan
		if !ok {
			break
		}else{
			for k, _ := range res{
				fmt.Println(k)
			}
		}

	}
}
