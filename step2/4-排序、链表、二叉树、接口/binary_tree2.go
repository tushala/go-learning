package main

import "fmt"

type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}
type BinTree struct {
	root *TreeNode
	left *TreeNode
	right *TreeNode
}
func (self *BinTree) PrintTree(p *TreeNode){
	if p == nil{
		return
	}
	self.PrintTree(p.left)
	fmt.Println(p.val)
	self.PrintTree(p.right)
}
func (self *TreeNode) String() string{
	s := "self.val"
	return s
}
func main() {
	left := TreeNode{
		val:   10,
		left:  nil,
		right: nil,
	}

	//right := TreeNode{
	//	val:   30,
	//	left:  nil,
	//	right: nil,
	//}
	//mid := TreeNode{
	//	val:   20,
	//	left:  &left,
	//	right: &right,
	//}
	//bt := BinTree{
	//	root:  &mid,
	//	left:  &left,
	//	right: &right,
	//}
	//printtree(&mid)
	//bt.PrintTree(&mid)
	fmt.Println(left)
}