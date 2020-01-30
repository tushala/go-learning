package main

import "fmt"

func badCall() {
	panic("bad end")
}
func test1() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicking %s\r\n", e)
		}
	}()
	badCall() //如果panic则停止不执行 fmt.Printf("After bad call\r\n")
	fmt.Printf("After bad call\r\n")
}

func main() {
	fmt.Printf("Calling test\r\n")
	test1()
	fmt.Printf("Test completed\r\n")
}
