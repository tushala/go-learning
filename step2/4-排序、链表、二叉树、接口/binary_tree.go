package main

import "fmt"

type TreeNode struct {
	val   float32
	left  *TreeNode
	right *TreeNode
}

func printtree(p *TreeNode) {
	if p == nil{
		return
	}
	printtree(p.left)
	fmt.Println(p.val)
	printtree(p.right)
}

func main() {
	left := TreeNode{
		val:   10,
		left:  nil,
		right: nil,
	}

	right := TreeNode{
		val:   30,
		left:  nil,
		right: nil,
	}
	mid := TreeNode{
		val:   20,
		left:  &left,
		right: &right,
	}
	printtree(&mid)
}
