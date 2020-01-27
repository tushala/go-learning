// 链表
package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float32
	Next  *Student
}

func insert(p *Student, n int) {
	for i := 0; i < n; i++ {
		stu := Student{
			Name:  fmt.Sprintf("stu%d", i+1),
			Age:   i + 20,
			Score: float32(100 - i),
			Next:  nil,
		}
		p.Next = &stu
		p = p.Next
	}
}
func addnode(p *Student, n int){
	new_node := Student{
		Name:  fmt.Sprintf("new_node%d",n),
		Age:   100,
		Score: 100,
		Next:  nil,
	}
	for i := 0; i < n-1; i++ {
		p = p.Next
	}
	if p.Next != nil{
		cur := p.Next
		p.Next = &new_node
		new_node.Next = cur
	}else{
		p.Next = &new_node
	}

}
func delnode(p *Student, n int) {
	for i := 0; i < n-1; i++ {
		p = p.Next
	}
	if p.Next != nil {
		p.Next = p.Next.Next
	}else{
		p.Next = nil
	}
}
func trans(p *Student) {
	for p != nil {
		fmt.Printf("%s -> ", p.Name)
		p = p.Next
	}
	fmt.Println()
}
func main() {
	var head = Student{
		Name:  "土沙拉",
		Age:   18,
		Score: 100,
		Next:  nil,
	}
	p := &head
	insert(p, 10)
	trans(p)
	delnode(p , 6)
	addnode(p, 9)
	trans(p)
}
