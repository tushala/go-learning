package main

import (
	"fmt"
	"go-learning/step2/流程控制、函数详解/utils"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func practice1(url string) {
	//判断一个url是否以http://开头，如果不是，则加上http://。
	//if strings.HasPrefix(url,"http://"){
	//	fmt.Println(url)
	//}else{
	//	fmt.Println("http://" + url)
	//}

	//判断一个路径是否以“/”结尾，如果不是，则加上/。
	if strings.HasPrefix(url, "/") {
		fmt.Println(url)
	} else {
		fmt.Println(url + "/")
	}
}
func practice2() {
	// Index首次出现 LastIndex最后出现 不存在返回-1
	fmt.Println(strings.Index("tushala", "a"))     //4
	fmt.Println(strings.Index("tushala", "x"))     //-1
	fmt.Println(strings.LastIndex("tushala", "a")) //6
	fmt.Println(strings.LastIndex("tushala", "x")) //-1
	// Replace、Count、Repeat、ToLower、ToUpper用法
	fmt.Println(strings.Replace("tusssshala", "s", "w", 2))
	fmt.Println(strings.Count("tusssshala", "s"))
	fmt.Println(strings.Repeat("tusssshala", 2))
	fmt.Println(strings.ToUpper("tusssshala"))
	fmt.Println(strings.ToLower("TUSSSSHALA"))
	//TrimSpace、Trim、TrimLeft、TrimRight
	fmt.Println(strings.TrimSpace("tushala   ")) // 去掉字符串首尾空白字符
	fmt.Println(strings.Trim("tushala", "tu"))   // 去掉字符串首尾cut字符
	fmt.Println(strings.TrimLeft("tushala", "tu"))
	fmt.Println(strings.TrimRight("tushala", "la"))
	//Field、Split、Join
	fmt.Printf("%q\n", strings.Fields("abc def ghi"))       //返回str空格分隔的所有子串的slice
	fmt.Printf("%q\n", strings.Split("ttattatta", "a"))     //split()
	fmt.Println(strings.Join([]string{"abc", "def"}, "\t")) //join()
	// 类型转换 strconv Itoa、.Atoi
	fmt.Println(strconv.Itoa(1123) + "3") //int->str
	i, err := strconv.Atoi("1123")        //str->int
	if err != nil {
		panic(err)
	}
	fmt.Println(i + 3)
}
func practice3() {
	//time包 使用
	now := time.Now()
	fmt.Println(now.Year(), now.Month(), now.Day(), now.Minute(), now.Second())
	fmt.Printf("%02d/%02d/%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	//格式化
	fmt.Println(now.Format("02/1/2006 15:04")) //一定要是2006/1/02 15:04这个时间
	fmt.Println(now.Format("2006/1/02 15:04"))
	fmt.Println(now.Format("2006/1/02"))
}
func practice4() {
	//写一个程序，获取当前时间，并格式化成 2017/06/15 08:05:00形式
	now := time.Now()
	fmt.Println(now.Format("2006/1/02 15:04:00"))
}
func practice5_test() {
	time.Sleep(time.Second)
}
func practice5() {
	//写一个程序，统计一段代码的执行耗时，单位精确到微秒
	start := time.Now().UnixNano()
	practice5_test()
	end := time.Now().UnixNano()
	fmt.Printf("%d us", end-start)
}
func practice6() {
	//写一个程序，获取一个变量的地址，并打印到终端
	var a int = 1
	fmt.Println(&a)
}
func practice7(a *int) {
	/*在main函数中调用这个函数，并把修改前后的值打印到终端，观察结果
	写一个函数，传入一个int类型的指针，并在函数中修改所指向的值。*/
	*a += 1
}
func practice8() {
	/*写一个程序，从终端读取输入，并转成整数，如果转成整数出错
	则输出 “can not convert to int”，并返回。否则输出该整数*/
	var s string
	fmt.Scanf("%s", &s)
	number, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("can not convert to int")
	} else {
		fmt.Println(number)
	}
}
func practice9() {
	/*猜数字，写一个程序，随机生成一个0到100的整数n，然后用户在终端，
	输入数字，如果和n相等，则提示用户猜对了。如果不相等，则提示用户，大于
	或小于n。*/
	fmt.Println("请输入一个0到100的整数")
	var input int
	rand.Seed(time.Now().UnixNano()) //设置随机种子

	ans := rand.Intn(100) //不设置随机种子，每次运行结果都一样
	flag := false
	for {
		fmt.Scanf("%d", &input)
		if input > 100 || input < 0 {
			panic("输入有误")
		}
		switch {
		case ans == input:
			fmt.Println("猜对了")
			flag = true
		case ans > input:
			fmt.Println("猜小了")
		default:
			fmt.Println("猜大了")
		}
		if flag {
			break
		}
		//fmt.Println(ans)
	}

}
func practice10() {
	/*打印
	A
	AA
	AAA
	AAAA
	AAAAA
	*/
	var A string = "A"
	for i := 0; i <= 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%s", A)
		}
		fmt.Println()

	}
}
func practice11(arg ...int) int {
	//写一个函数add，支持1个或多个int相加，并返回相加结果
	sum := 0
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum
}
func practice12(arg ...string) string {
	// 写一个函数concat，支持1个或多个string相拼接，并返回结果
	return strings.Join(arg, "")
}
func homework1() {
	//编写程序，在终端输出九九乘法表。
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = *%d\t\t", i, j, i*j)
		}
		fmt.Println()
	}
	fmt.Println()
}
func homework2() {
	//一个数如果恰好等于它的因子之和，这个数就称为“完数”。例如6=1＋2＋3.
	//编程找出1000以内的所有完数。
	for i := 2; i < 1000; i++ {
		if utils.Is_WanNum(i) {
			fmt.Println(i)
		}
	}

}
func homework3(str string) {
	/*输入一个字符串，判断其是否为回文。回文字符串是指从左到右读和从右到
	左读完全相同的字符串。*/
	if utils.Is_Palindrome(str) {
		fmt.Printf("%s是回文数", str)
	} else {
		fmt.Printf("%s不是回文数", str)
	}
}
func homework4() {
	/*输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数*/
	var input string
	var char = 0
	var space = 0
	var num = 0
	var other = 0
	fmt.Scanf("%s", &input)
	fmt.Printf("输入字符为: %s\n", input) //todo 空格不能计入？？
	for _, v := range []byte(input) {
		fmt.Println(v)
		switch {
		case (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z'):
			char += 1

		case v == ' ':
			space += 1

		case v >= '0' && v <= '9':
			num += 1
		default:
			other += 1
		}
	}
	fmt.Printf("英文字母%d个、空格%d个、数字%d个和其它字符%d个", char, space, num, other)
}
func homework5(a, b int) {
	//计算两个大数相加的和，这两个大数会超过int64的表示范围.
	stra := strconv.Itoa(a)
	strb := strconv.Itoa(b)
	borrow := 0
	var ans string
	var s int

	min_length := int(math.Min(float64(len(stra)), float64(len(strb))))
	if len(strb) >= len(stra) {
		stra, strb = strb, stra
	}
	length_a := len(stra)
	length_b := len(strb)
	for i := 0; i < min_length; i++ {
		va := int(stra[length_a-i-1] - '0')
		vb := int(strb[length_b-i-1] - '0')
		if va+vb >= 10 {
			s = va + vb + borrow - 10
			borrow = 1
		} else {
			s = va + vb + borrow
			borrow = 0
		}

		p := string('0' + s)
		ans = p + ans
	}

	//for i := len(stra)-1; i <= min_length-1; i++ {
	for i := min_length; i < len(stra); i++ {
		//fmt.Println(123, i)
		s := int(stra[length_a-i-1] - '0')
		if s+borrow >= 10 {
			s = s + borrow - 10
			borrow = 1
		} else {
			s = s + borrow
			borrow = 0
		}
		p := string('0' + s)
		ans = p + ans
	}
	if borrow == 1 {
		ans = "1" + ans
	}
	fmt.Println(ans)
}
func main() {
	//strings和strconv使用
	//practice1("http://abc/")
	//practice2()
	//practice3()
	//practice4()
	//practice5()
	//practice6()

	//var a int = 1
	//fmt.Println(a)
	//practice7(&a)
	//fmt.Println(a)
	//practice8()

	//var a, b, c int
	//for i := 0; i < 2; i++ {
	//	fmt.Scanf("%d,%d,%d", &a, &b, &c)
	//	fmt.Println(a, b, c)
	//}
	//practice9()
	//practice10()

	//f := practice10 //一个函数可以赋值给变量
	//f()

	//fmt.Println(practice11(1,2,3,4,5,6))
	//fmt.Println(practice12("a","b","c","d"))
	//homework1()
	//homework2()
	//homework3("abcaaa")
	//homework4()
	homework5(1, 99999999)
}
