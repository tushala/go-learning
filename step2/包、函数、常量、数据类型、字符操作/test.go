package main

import "fmt"

var a = "G"


func n(){
	fmt.Println(a)
}

func m(){
	a := "O"
	fmt.Println(a)
}
func main() {
	//n()
	//m()
	//n()
	//var n int16=34
	//var m int32
	//
	////m = n
	//m =int32(n)
	//fmt.Printf("32 bit int is: %d\n", m)
	//fmt.Printf("16 bit int is: %d\n", n)
	//var a byte = 'a' //单引号
	//var str = "hello world"
	//fmt.Println("a = ", a)  //97
	//fmt.Println("str = ", str)
	const(
		a = iota
		b
		c
	)
	fmt.Println(a, b, c) // 0 1 2
}