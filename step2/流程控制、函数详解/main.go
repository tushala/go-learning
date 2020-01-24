package main

import (
	"fmt"
	"strings"
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
func practice2(){
	// Index首次出现 LastIndex最后出现 不存在返回-1
	fmt.Println(strings.Index("tushala", "a"))     //4
	fmt.Println(strings.Index("tushala", "x"))     //-1
	fmt.Println(strings.LastIndex("tushala", "a")) //6
	fmt.Println(strings.LastIndex("tushala", "x")) //-1
	// Replace、Count、Repeat、ToLower、ToUpper用法
	fmt.Println(strings.Replace("tusssshala", "s", "w", 2))
	fmt.Println(strings.Count("tusssshala", "s"))
	fmt.Println(strings.Repeat("tusssshala",2))
	fmt.Println(strings.ToUpper("tusssshala"))
	fmt.Println(strings.ToLower("TUSSSSHALA"))
}
func main() {
	//strings和strconv使用
	//practice1("http://abc/")
	practice2()


}
