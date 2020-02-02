package main

import "fmt"

type Person struct {
	Name, Address string
	Age           int
}

func main() {
	var c chan interface{}
	var a Person
	c = make(chan interface{}, 10)
	for i := 0; i < 10; i++ {
		per := Person{
			Name:    fmt.Sprintf("stu%d", i+1),
			Address: "abcde",
			Age:     20 + i,
		}
		c <- per
	}
	for i := 0; i < 10; i++{
		new_person := <-c  //new_person默认inerface{}类型
		a  = new_person.(Person)
		fmt.Println(a.Name)
	}

}
