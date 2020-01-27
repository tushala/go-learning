package main

import "time"

func test1(){
	// 匿名字段
	type Car struct {
		Name string
		Age int
	}
	type Train struct{
		Car
		Start time.Time
		int
	}
}
func main() {

}
