package main

import "fmt"

type LinkNode struct {
	data int
	next *LinkNode
}
type Link struct {
	head *LinkNode
	tail *LinkNode
}

func (self *Link) InsertHead(v int) {
	node := &LinkNode{
		data: v,
		next: nil,
	}
	if self.tail == nil && self.head == nil {
		self.head = node
		self.tail = node
		return
	}
	node.next = self.head
	self.head = node
}
func (self *Link) InsertTail(v int) {
	node := &LinkNode{
		data: v,
		next: nil,
	}
	if self.tail == nil && self.head == nil {
		self.head = node
		self.tail = node
		return
	}
	self.tail.next = node
	self.tail = node
}
func (self *Link) Trans() {
	p := self.head
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}
}
func main() {
	var link Link
	for i := 0; i < 10; i++ {
		link.InsertHead(i)
		link.InsertTail(i)
	}
	link.Trans()
}
