package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
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
		case v >= '0' && v <= '9':
			num += 1
		default:
			other += 1
		}
	}
	fmt.Printf("char %d num %d other %d\n", char, num, other)
}
func practice5(filepath string) {
	// 文件操作示例
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("read file err: ", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读一行\n分割
		if err == io.EOF {                  // EOF文件读取完毕
			break
		} else if err != nil {
			fmt.Println("read string failed, err: ", err)
			return
		}
		fmt.Printf("%s ", str)
	}
}
func practice6(inputFile, outputFile string) {
	// 读取整个文件示例
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		return
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0x644)
	if err != nil {
		panic(err.Error())
	}
}
func practice7(inputFile, outputFile string)(written int64, err error){
	src, err := os.Open(inputFile)
	if err != nil{
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil{
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)

}
func practice8(){
	// os使用
	var(
		test bool
		str string
		count int
	)
	flag.BoolVar(&test, "b", false, "print on newline")
	flag.StringVar(&str, "s","","print on newline")
	flag.IntVar(&count, "c", 1001, "print on newline")
	flag.Parse()

	fmt.Println(test)
	fmt.Println(str)
	fmt.Println(count)
}
func practice9(){
	// json
	// golang ->(序列化)json字符串 ->(网络传输)程序->(反序列化)->其他语言
	// struct 序列化
	s := &Student{
		Name:  "tushala",
		Age:   26,
		Score: 100,
	}
	data, err := json.Marshal(s)
	if err != nil{
		fmt.Printf("json marshal failed, err: ", err)
		return
	}
	fmt.Printf("%s\n", string(data))
	// 反序列化
	var s1 Student
	err = json.Unmarshal([]byte(data), &s1)
	if err != nil{
		fmt.Printf("json Unmarshal failed, err: ", err)
		return
	}
	fmt.Println(s1)
}
func practice10(){
	/*map 序列化
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["a"] = "kobe"
	m["b"] = 41
	m["c"] = 100.0
	data, err := json.Marshal(m)
	if err != nil{
		fmt.Printf("json marshal failed, err: ", err)
		return
	}
	fmt.Printf("%s\n", string(data))*/

	//
	var m []interface{}
	//m = make([]interface{}, 3)
	m = append(m, "kobe")
	m = append(m, "81")
	m = append(m, "41")
	data, err := json.Marshal(m)
	if err != nil{
		fmt.Printf("json marshal failed, err: ", err)
		return
	}
	fmt.Printf("%s\n", string(data))
}

func main() {
	//practice1()
	//practice2()
	//practice3()
	//practice4()
	//practice5("C:/Users/19416/Desktop/导数据库.txt")
	//practice6("C:/Users/19416/Desktop/导数据库.txt", "1.txt") //在src目录下
	//practice7("C:/Users/19416/Desktop/导数据库.txt", "1.txt")
	//practice8()
	//practice9()
	practice10()
}
