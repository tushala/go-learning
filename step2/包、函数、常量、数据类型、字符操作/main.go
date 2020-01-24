package main

import (
	// 从外引入需大写
	"fmt"
	_ "go-learning/step2/包、函数、常量、数据类型、字符操作/add" //_ 导入但不用
	"math/rand"
	"os"
	"time"
	//"./add"
	"./utils"
)

func init() {
	/*每个源文件都可以包含一个init函数
	这个init函数自动被go运行框架调用*/
	fmt.Println("每个源文件都可以包含一个init函数，这个init函数自动被go运行框架调用")
}
func practice1(num int) {
	//写一个程序，对于给定一个数字n，求出所有两两相加等于n的组合
	for i := 0; i <= num; i++ {
		fmt.Printf("%d + %d = %d\n", i, num-i, num)
	}
}

//func practice2(){
//	fmt.Println(add.Name)
//	fmt.Println(add.Age)
//}
//func practice3(){
//	// 包别名的应用
//	fmt.Println(a.Name)
//	fmt.Println(a.Age)
//}
func practice4() {
	/*定义两个常量Man=1和Female=2，获取当前时间的秒数，如果能被Female整除，则
	在终端打印female，否则打印man。*///ctrl+shift+/

	const Man = 1
	const Female = 2
	Second := time.Now().Unix()
	if Second%Female == 1 {
		fmt.Println("female")
	} else {
		fmt.Println("male")
	}
}
func practice5() {
	/*写一个程序获取当前运行的操作系统名称和PATH环境环境变量的值，并打印在终端。*/
	var goos string = os.Getenv("PATH")
	fmt.Println(goos)
}
func practice6() {
	//写一个程序用来打印值类型和引用类型变量到终端，并观察输出结果
	/*值类型: 基本数据类型,数组和struct
	引用类型 : 指针、slice、map、chan*/
	var a = 100
	var b chan int = make(chan int, 1)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
func swap(a, b *int) {
	*a, *b = *b, *a
}
func practice7() {
	//写一个程序，交换两个整数的值
	var a = 3
	var b = 4
	fmt.Printf("a = %d b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("a = %d b = %d", a, b)
}
func practice8() {
	var a bool = true
	var b bool
	fmt.Println(!a)
	fmt.Println(!b)
	fmt.Println(a && b)
	fmt.Println(a || b)
}
func practice9() {
	//类型转换
	var a int = 15
	var b int32 = int32(a)
	//b = int32(a)
	fmt.Println(b)
}
func practice10() {
	//使用math/rand生成10个随机整数，10个小于100的随机整数以及10个随机浮点数？
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", rand.Intn(100))
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Printf("%.2f ", 100*rand.Float32())
	}
}
func practice11() {
	//字符串表示两种方式： 1）双引号    2）``   （反引号）
	var str1 = "hello world"
	var str2 = `hello world\n\n\n\
tutututushala
tututshala`
	fmt.Println("str1=", str1)
	fmt.Println("str2=", str2)
}
func homework1() {
	//判断 101-200 之间有多少个素数，并输出所有素数
	for i := 101; i < 201; i++ {
		if utils.Is_Prime(i) == true {
			fmt.Println(i)
		}
	}
}
func homework2() {
	//打印出100-999中所有的“水仙花数”，
	for i := 101; i < 1000; i++ {
		if utils.Is_Narcissus_Num(i) == true {
			fmt.Println(i)
		}
	}
}
func homework3(num int){
	fmt.Println(utils.Factorial_sum(num))
}
func main() {
	//practice1(5)
	//practice2()
	//practice3()
	//practice4()
	//practice5()
	//practice6()
	//practice7()
	//practice8()
	//practice9()
	//practice10()
	//practice11()
	//homework1()
	//homework2()
	homework3(5)
}
