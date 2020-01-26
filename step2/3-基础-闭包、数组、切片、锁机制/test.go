package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

func test1() {
	// 数组初始化
	var age0 [5]int = [5]int{1, 2, 3}
	fmt.Println(age0)
	var age1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(age1)
	var age2 = [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(age2)
	var str = [5]string{3: "hello world", 4: "tom"} // 第3个为hello world，第四个为tom
	fmt.Println(str)

	// 多维数组
	var age [5][3]int
	var f [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(age)
	fmt.Println(f)
}
func test2() {
	var f [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for k1, v1 := range f {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d\t", k1, k2, v2)
		}
		fmt.Println()
	}
}
func test3() {
	// 切片
	var slice []int
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	//slice = arr[2:3]
	slice = arr[:]
	fmt.Println(slice)
	//fmt.Println(len(slice)) // 长度3
	//fmt.Println(cap(slice)) // 容量3
	slice = slice[0:1]
	fmt.Println(len(slice)) // 长度1
	fmt.Println(cap(slice)) // 容量3
}
func test4() {
	// 因此切片是引用类型
	var arr = [...]int{1, 2, 3}
	s := arr[:]
	arr[1] = 100

	s[0] = 200
	fmt.Println(arr)
	fmt.Println(s)
}
func test5() {
	// 切片的内存布局 验证数组的第i个地址和对应的切片地址相同
	var a = [3]int{1, 2, 3}
	s := a[:]
	fmt.Println(&a[0])
	fmt.Println(&s[0])
	fmt.Printf("%p\n", s) // 对应切片头地址
}
func test6() {
	//make创建切片
	/*var s []type = make([]type, 5)
	s1 := make([]type, 5)
	s2 := make([]types, 5, 10)
	s1 := make([]int, 5)
	s1 = []int{1,2,3,4,5}
	fmt.Println(s1)*/
	//切片拷贝
	/*s1 := []int{1,2,3,4,5}
	s2 := make([]int, 10)
	copy(s2, s1)
	fmt.Println(s2) //[1 2 3 4 5 0 0 0 0 0]
	s3 := []int{1,2,3}
	s3 = append(s3, s2...)
	fmt.Println(s3)
	s3 = append(s3, 4,5,6)
	fmt.Println(s3)*/
	// string底层就是一个byte的数组，因此，也可以进行切片操作
	//str := "hello world"
	//s1 := str[:5]
	//fmt.Println(s1)
	//s2 := str[5:]
	//fmt.Println(s2)

	//for _, v := range str{
	//	fmt.Println(string(v))
	//}

	//s := []byte(str)
	//s[0] = 'O'
	//fmt.Println(string(s))
	var a = [...]int{5, 4, 3, 2, 1}
	sort.Ints(a[:])
	fmt.Println(a)
}
func modify(a map[string]string){
	a["hello"] = "tushala"
}
func test7() {
	// map
	//map赋值前要先初始化
	//var a map[string]string = map[string]string{"hello": "world"}
	//a := make(map[string]string, 5)
	//a["hello"] = "world"
	//a["hello1"] = "world1"
	//a["hello2"] = "world2"
	//val, ok := a["hello"]
	//fmt.Println(val, ok)
	//
	//for k, v := range a{
	//	fmt.Println(k, v)
	//}
	//fmt.Println(len(a))
	//delete(a, "hello")
	//val, ok = a["hello"]
	//fmt.Println(val, ok)
	//fmt.Println(len(a))
	//map是引用类型
	/*a := make(map[string]string, 5)
	a["hello"] = "world"
	a["hello1"] = "world1"
	modify(a)
	fmt.Println(a)*/
	//slice of map
	items := make([]map[int]int, 5)
	for i:=0;i<5;i++{
		items[i] = make(map[int]int)
	}
}
func test8(){
	// map排序
	var a = make(map[int]int, 5)
	var keys []int
	a[8] = 10
	a[6] = 20
	a[4] = 30
	a[2] = 40
	a[0] = 50
	for k, _ := range a{
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, v := range keys{
		fmt.Println(v, a[v])
	}
}
func test9(){
	// Map反转
	var a = make(map[int]int, 5)
	var r = make(map[int]int, 5)

	a[8] = 10
	a[6] = 20
	a[4] = 30
	a[2] = 40
	a[0] = 50
	for k, v := range a{
		r[v] = k
	}
	fmt.Println(r)
}
func test10(){
	// 锁示例
	var lock sync.Mutex
	var a = make(map[int]int, 5)
	a[8] = 10
	a[6] = 20
	a[4] = 30
	a[2] = 40
	a[0] = 50
	for i:=0;i<2;i++{
		go func(b map[int]int){
			lock.Lock()
			b[0] = rand.Intn(200)
			lock.Unlock()
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()
}
func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//test7()
	//test8()
	//test9()
	test10()
}
