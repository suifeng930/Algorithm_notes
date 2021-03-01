package main

import "log"

var i int


func main() {
	i =-1
	arrs :=[]interface{}{3,2,9,nil,nil,10,nil,nil,8,nil,4}
	treeNode := createBinaryTree(arrs)
	//PrintNode(treeNode)
	log.Println("前序遍历")
	preOrderTraveral(treeNode)
	log.Println("中序遍历")
	inOrderTraveral(treeNode)
	log.Println("后序遍历")
	postOrderTraveral(treeNode)


}

type TreeNode struct {
	Value      interface{}
	LeftChild  *TreeNode
	RightChild *TreeNode
}

//前序遍历 先根-->再左子树-->最后右子树
func preOrderTraveral(node *TreeNode) {

	if node == nil {
		return
	}
	log.Println(node.Value)
	preOrderTraveral(node.LeftChild)
	preOrderTraveral(node.RightChild)

}

// 中序遍历，先左子树-->再根-->最后右子树
func inOrderTraveral(node *TreeNode) {
	if node == nil {
		return
	}
	inOrderTraveral(node.LeftChild)
	log.Println(node.Value)
	inOrderTraveral(node.RightChild)
}

// 后序遍历，先左子树-->再右子树-->最后根
func postOrderTraveral(node *TreeNode) {
	if node == nil {
		return
	}
	postOrderTraveral(node.LeftChild)
	postOrderTraveral(node.RightChild)
	log.Println(node.Value)
}

//创建一个二叉树
func createBinaryTree(arr []interface{}) *TreeNode {

	i = i + 1
	if i>=len(arr){
		return nil
	}
	var node TreeNode
	if arr[i] != 0 {
		node = TreeNode{Value: arr[i]}
		node.LeftChild = createBinaryTree(arr)
		node.RightChild = createBinaryTree(arr)
	} else {
		return nil
	}
	return &node

}

func PrintNode(node *TreeNode)  {
	if node==nil {
		return
	}
	log.Printf("%v",node.Value)

}