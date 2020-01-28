// sort 举例
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name string
	Id string
	Age int
}
type StudentArray []Student
func (self StudentArray) Len() int{
	return len(self)
}
func (self StudentArray) Less(i, j int) bool{
	return self[i].Name > self[j].Name
}
func (self StudentArray) Swap(i,j int){
	self[i], self[j] = self[j], self[i]
}
func main() {
	var stus StudentArray
	for i:=0;i<10;i++{
		stu := Student{
			Name: fmt.Sprintf("stu%d", rand.Intn(100)),
			Id:   fmt.Sprintf("110%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		stus = append(stus, stu)
	}
	for _, v := range stus{
		fmt.Println(v)
	}
	fmt.Println("\n\n")
	sort.Sort(stus)
	for _, v := range stus{
		fmt.Println(v)
	}
}