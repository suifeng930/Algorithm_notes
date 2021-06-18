package utils

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root  *TreeNode)PrintTree()  {

	stack :=[]*TreeNode{}
	cur :=root
	for cur!=nil ||len(stack)>0 {

		for cur!=nil {
			stack=append(stack,cur)
			cur=cur.Left
		}
		cur=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		fmt.Printf("%v-->",cur.Val)
		cur=cur.Right
	}
	fmt.Println()
}


func NewNode(v int) *TreeNode {
	return &TreeNode{
		Val:      v,
	}
}