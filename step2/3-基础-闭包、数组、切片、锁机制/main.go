package main

import (
	"fmt"
	"strings"
)

func practice1() {
	// new：用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针, 返回儲存位址
	var i int
	fmt.Println(i)
	j := new(int)
	*j = 100
	fmt.Println(*j)
}
func practice2() {
	// make：用来分配内存，主要用来分配引用类型，比如chan、map、slice make 不会回传地址
	var i map[string]int
	fmt.Println(i)
	j := make(map[string]int)
	fmt.Println(j)
}
func practice3() {
	//append：用来追加元素到数组、slice中
	var a []int
	a = append(a, 10, 20, 30)
	fmt.Println(a)
	a = append(a, a...) // ...表示展开
	fmt.Println(a)
}

//func practice4() {
//	// panic和recover：用来做错误处理
//	defer func(){
//		if err := recover(); err !=nil{
//			fmt.Println(err)
//		}
//	}()
//	b := 0
//	a := 100 / b
//	fmt.Println(a)
//}
func practice5(n int) int {
	// 递归
	if n == 1 {
		return 1
	}
	return practice5(n-1) * n
}
func practice6(n int) int {
	if n <= 1 {
		return 1
	}
	return practice6(n-1) + practice6(n-2)
}
func practice7() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
func practice8(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func practice9(arr [5]int) {
	// 数组是值类型，因此改变副本的值，不会改变本身的值
	arr[0] = 100
	return
}
func practice10() {
	//使用非递归的方式实现斐波那契数列，打印前50个数。
	var a uint64 = 1
	var b uint64 = 1
	//a, b := 1, 1
	for i := 0; i < 50; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}
func homework1(a []int) {
	//冒泡排序
	for i:=0;i<len(a);i++{
		for j:=0;j<i;j++{
			if a[i] < a[j]{
				a[i],a[j] = a[j], a[i]
			}
		}
	}
}
func main() {
	//practice1()
	//practice2()
	//practice3()
	/*for{
		practice4()
	}*/
	/*f := practice7()
	fmt.Println(f(1))   // 1
	fmt.Println(f(20))  // 21
	fmt.Println(f(300)) //321*/

	/*f1 := practice8(".bmp")
	f2 := practice8(".jpg")
	fmt.Println(f1("test"))
	fmt.Println(f2("test"))*/

	/*var a [5]int
	practice9(a)
	for i:=0;i<len(a);i++{
		fmt.Println(a[i])
	}*/
	//practice10()
	var a = [...]int{6, 5, 4, 3, 2, 1}
	fmt.Println(a)
	homework1(a[:])
	fmt.Println(a)
}
