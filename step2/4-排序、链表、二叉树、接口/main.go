package main

import (
	"encoding/json"
	"fmt"
)

func practice1() {
	type Student struct {
		Name  string // this is Name Field  tag
		Age   int
		Score float32
	}
	//var stu Student
	//var stu *Student = new(Student)
	var stu *Student = &Student{}
	stu.Name = "tony"
	stu.Age = 21
	stu.Score = 99.532
	//fmt.Printf("name=%s age=%d score=%.2f",
	//	stu.Name, stu.Age, stu.Score)
	fmt.Printf("name=%s age=%d score=%.2f",
		(*stu).Name, (*stu).Age, (*stu).Score)
}
func practice2() {
	type intrgar int
	var i intrgar = 1 // 别名
	fmt.Println(i)
}
func practice3() {
	// json序列化和反序列化
	type Student struct {
		Name string `json:"this is Name Field"`
		Age  int    `json:"this is Age Field"`
	}

	stu := Student{
		Name: "tushala",
		Age:  26,
	}
	data, err := json.Marshal(stu)
	if err != nil {
		return
	}
	fmt.Println(string(data))
}
func practice4(name string, age int) {
	// 工厂模式创建新的实例
	//return &St{
	//	Name: name
	//	Age:age
	//}
}
func practice5(items ...interface{}) {
	fmt.Println(items)
	//写一个函数判断传入参数的类型
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float32:
			fmt.Printf("param #%d is a float32\n", i)
		case int:
			fmt.Printf("param #%d is a int\n", i)
		case nil:
			fmt.Printf("param #%d is a nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d’s type is unknown\n", i)
		}
	}
}
func practice6() {
	var x int
	var i interface{}
	i = x
	//y, ok := i.(int64)
	y, ok := i.(int)
	fmt.Println(y, ok)
}
func main() {
	//practice1()
	//practice2()
	//practice3()
	//type St struct {
	//	Name string
	//	Age int
	//}
	//practice4()

	//practice5("tushala", 123, true, nil, 1.2)
	practice6()

}
