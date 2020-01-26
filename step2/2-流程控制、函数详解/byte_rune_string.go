// byte、rune与string 学习
package main

import "fmt"

func main() {
	//1. byte 与 rune
	//双引号，表示字符串，而在单引号表示rune 类型的字符
	//byte，占用1个节字，就 8 个比特位，所以它和 uint8 类型本质上没有区别，它表示的是 ACSII 表中的一个字符。
	//var a byte = 65
	//var b uint8 = 66
	//fmt.Println("a 的值: ", string(a)," \nb 的值: ", string(b))
	//fmt.Printf("a 的值: %c \nb 的值: %c", a, b)

	//var a byte = 'A'
	////byte 类型能表示的值是有限，只有 2^8=256 个。所以如果你想表示中文的话，你只能使用 rune 类型。
	//var b rune = 'B'  // rune，占用4个字节，共32位比特位
	//fmt.Printf("a 占用 %d 个字节数\nb 占用 %d 个字节数\n", unsafe.Sizeof(a), unsafe.Sizeof(b))
	//var name rune = '中'
	//fmt.Println(name, string(name))

	//2. 字符串
	// 多个字符放在一起，就组成了字符串
	var mystr1 string = "hello"
	var mystr2 [5]byte = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("mystr1: %s\n", mystr1)
	fmt.Printf("mystr2: %s\n", mystr2)

	var mystr3 string = "hello,中国"
	var mystr4 [8]rune = [8]rune{104, 101, 108, 108, 111, 44, 20013, 22269}

	fmt.Printf("mystr3: %s\n", mystr3)
	fmt.Printf("mystr4: %s\n", mystr4)
	for _, v := range mystr4{
		fmt.Printf("%c ",v)
	}
}
