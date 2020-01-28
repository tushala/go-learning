package main

import "fmt"

/*实现一个图书管理系统，具有以下功能：
a. 书籍录入功能，书籍信息包括书名、副本数、作者、出版日期
b. 书籍查询功能，按照书名、作者、出版日期等条件检索
c. 学生信息管理功能，管理每个学生的姓名、年级、身份证、性别、借了什么书等信息
d. 借书功能，学生可以查询想要的书籍，进行借出
e. 书籍管理功能，可以看到每种书被哪些人借出了*/

func f1(a int)(ans bool){
	if a > 7{
		ans = true
		return
	}
	return
}
func main() {
	fmt.Println(f1(5))
	fmt.Println(f1(10))
}