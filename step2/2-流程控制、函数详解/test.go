package main

import "fmt"

func test1() {
	//case 用法1，fallthrough
	var i = 0
	switch i {
	case 0, 1:
		fmt.Println(1)
		fallthrough //fallthrough不会判断下一条case的expr结果是否为true
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("def")
	}
}
func test2() {
	i := 0
	switch {
	case i > 0 && i < 10:
		fmt.Println("i > 0 and i < 10")
	case i > 10 && i < 100:
		fmt.Println("i > 10 and i < 100")
	default:
		fmt.Println("def")
	}
}
func test3() {
	str := "hello world,中国"
	for i, v := range str {
		fmt.Printf("index[%d] val[%c] \n", i, v)
	}
}
func test4() {
	str := "hello world,中国"
	for i, v := range str {
		if i > 2 {
			continue
		}
		if (i > 3) {
			break
		}
		fmt.Printf("index[%d] val[%c]\n", i, v, )
	}
}
func test5() {
	//goto label
}
func test8(a, b, c int) int {
	return a + b
}
func test9(a *int) {
	*a = 100
}

func test10(a, b int) (c int) {
	c = a + b
	return
}
func test11(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a + b) / 2
	return
}
func test12(a,b int , arg ...int)int{
	// ...int类似
	fmt.Println(a, b)
	fmt.Println(arg)
	return arg[0]
}
func test13(){
	// 栈操作
	i := 0
	defer fmt.Println(i)
	defer fmt.Println("tushala")
	i+=1
	return
}
func test14()int{
	//43210100
	for i := 0; i < 5; i++{
		defer fmt.Printf("%d", i)
	}
	return 100
}
//defer用途：
//1.关闭文件句柄
//2.锁资源释放
//3.数据库连接释放
func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//test8(2,3,5)

	//a := 8
	//test9(&a)
	//fmt.Println(a)
	//fmt.Println(test11(2, 4))
	//fmt.Println(test12(1,2,3,4))
	//test13()
	fmt.Println(test14())
}
