package main

import (
	"fmt"
	"sync"
)

func Productor(mychan chan int, data int, wg *sync.WaitGroup) {
	mychan <- data
	fmt.Println("product data：", data)
	wg.Done()
}
func Consumer(mychan chan int, wg *sync.WaitGroup) {
	a := <-mychan
	fmt.Println("consumer data：",a)
	wg.Done()
}
func main() {
	datachan := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go Productor(datachan, i, &wg)
		wg.Add(1)
	}
	for j := 0; j < 10; j++ {
		go Consumer(datachan, &wg)
		wg.Add(1)
	}
	wg.Wait()

}
