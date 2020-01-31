package main

import (
	"fmt"
	"time"
)

func sendData(ch chan string){
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string){
	var input string
	for {
		input = <- ch
		fmt.Println(input)
	}
}

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
	time.Sleep(10*time.Second)
}