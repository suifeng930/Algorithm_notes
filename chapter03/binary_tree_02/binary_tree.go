package main

import (
	"log"
)

func main() {

	// 构造二叉树
	root := NewNode(3)
	root.LeftChild = &TreeNode{}
	root.LeftChild.setValue(2)
	root.LeftChild.LeftChild = NewNode(9)
	root.LeftChild.RightChild = NewNode(10)
	root.RightChild = NewNode(8)
	root.RightChild.RightChild = NewNode(4)

	log.Println("前序遍历")
	preOrderTraveral(root)
	log.Println("中序遍历")
	inOrderTraveral(root)
	log.Println("后序遍历")
	postOrderTraveral(root)
	log.Println("层序遍历")
	levelOrderTraveral(root)

}

type TreeNode struct {
	Value      interface{}
	LeftChild  *TreeNode
	RightChild *TreeNode
}

//前序遍历 每到一个节点A，就应该立即访问它。因为每棵树都是先访问其根节点。
//对节点的左右子树来说，也一定是先访问其根节点。在A的两颗子树中，遍历完左子树后，再遍历右子树。
//因此在访问完根节点后，遍历左子树前，要将右子树压入栈中。
func preOrderTraveral(root *TreeNode) {

	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			log.Println(node.Value)
			stack = append(stack, node)
			node = node.LeftChild
		}
		node = stack[len(stack)-1].RightChild //将右孩子压入栈中
		stack = stack[:len(stack)-1] //出栈
	}

}

// 中序遍历，先左子树-->再根-->最后右子树
func inOrderTraveral(root *TreeNode) {
	stack :=[]*TreeNode{}
	for root!=nil ||len(stack)>0 {
		for root!=nil {
			stack=append(stack,root)
			root=root.LeftChild
		}
		root=stack[len(stack)-1]
		stack=stack[:len(stack)-1]//出栈
		log.Println(root.Value)
		root=root.RightChild
	}

}

// 后序遍历，先左子树-->再右子树-->最后根
func postOrderTraveral(root *TreeNode) {
	stack :=[]*TreeNode{}
	var prev *TreeNode
	for root!=nil ||len(stack)>0 {
		for root!=nil {
			stack=append(stack,root)
			root=root.LeftChild
		}
		root=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		if root.RightChild==nil ||root.RightChild==prev {
			log.Println(root.Value)
			prev=root
			root=nil
		}else {
			stack=append(stack,root)
			root=root.RightChild
		}
	}
}

func levelOrderTraveral(root *TreeNode) {
	ret := [][]interface{}{} //二维数组记录层序以及每层都节点数量
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ { // 遍历层序
		ret = append(ret, []interface{}{})
		temp := []*TreeNode{}
		for j := 0; j < len(queue); j++ { //遍历每层都节点数量
			node := queue[j]
			log.Println(node.Value)
			ret[i] = append(ret[i], node.Value)
			if node.LeftChild != nil {
				temp = append(temp, node.LeftChild)
			}
			if node.RightChild != nil {
				temp = append(temp, node.RightChild)
			}
		}
		queue = temp
	}
	log.Println("ret:", ret)
}

func PrintNode(node *TreeNode) {
	if node == nil {
		return
	}
	log.Printf("%v", node.Value)

}
func (node *TreeNode) setValue(v interface{}) {
	if node == nil {
		log.Println("setting value  to nil node")
		return
	}
	node.Value = v
}

func NewNode(v interface{}) *TreeNode {
	return &TreeNode{
		Value: v,
	}

}
