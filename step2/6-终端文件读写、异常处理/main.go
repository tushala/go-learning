package main

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func practice1() {
	var str = "stu01 18 99.99"
	var stu Student
	fmt.Sscanf(str, "%s %d %f", &stu.Name, &stu.Age, &stu.Score)
	fmt.Println(stu)
}
func practice2() {
	var (
		firstName, lastName, s string
		i                      int
		f                      float32
		input                  = "56.12 / 5212 / Go"
		format                 = "%f / %d / %s"
	)
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	//fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)

}
func practice3() {
	// 带缓冲区的读写
	var inputReader *bufio.Reader
	var input string
	var err error

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
	fmt.Println(input)
}
func practice4() {
	var input string
	var char, num, other int
	fmt.Println("input: ")
	fmt.Scanf("%s", &input)
	for _, v := range input {
		switch {
		case (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z'):
			char += 1
		case (v >= '0' && v <= '9'):
			num += 1
		default:
			other += 1
		}
	}
	fmt.Printf("char %d num %d other %d\n", char, num, other)
}
func main() {
	//practice1()
	//practice2()
	//practice3()
	practice4()
}
