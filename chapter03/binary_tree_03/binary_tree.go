package main

import "log"

var i int


func main() {
	i =-1
	arrs :=[]interface{}{3,2,9,nil,nil,10,nil,nil,8,nil,4}
	treeNode := createBinaryTree(arrs)

	log.Println(treeNode)
	// 构造二叉树
	root := NewNode(3)
	root.LeftChild=&TreeNode{}
	root.LeftChild.setValue(2)
	root.LeftChild.LeftChild=NewNode(9)
	root.LeftChild.RightChild=NewNode(10)
	root.RightChild=NewNode(8)
	root.RightChild.RightChild=NewNode(4)


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
func inOrderTraveral(node *TreeNode){
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

// 二叉树的层序遍历   先将根节点入队列，遍历队列，判断节点是否存在根节点
func levelOrderTraveral(root *TreeNode) {
	ret :=[][]interface{}{} //二维数组记录层序以及每层都节点数量
	if root ==nil {
		return
	}
	queue :=[]*TreeNode{root}
	for i:=0;len(queue)>0;i++ { // 遍历层序
		ret =append(ret,[]interface{}{})
		temp :=[]*TreeNode{}
		for j :=0;j<len(queue);j++ { //遍历每层都节点数量
			node :=queue[j]
			log.Println(node.Value)
			ret[i]=append(ret[i],node.Value)
			if node.LeftChild!=nil{
				temp=append(temp,node.LeftChild)
			}
			if node.RightChild!=nil {
				temp=append(temp,node.RightChild)
			}
		}
		queue=temp
	}
	log.Println("ret:",ret)
}

//创建一个二叉树
func createBinaryTree(arr []interface{}) *TreeNode {

	i = i + 1
	if i>=len(arr){
		return nil
	}
	var node TreeNode
	if arr[i] != 0 {
		node = TreeNode{arr[i],nil,nil}
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
func (node *TreeNode) setValue(v interface{})  {
	if node ==nil {
		log.Println("setting value  to nil node")
		return
	}
	node.Value=v
}

func NewNode(v interface{}) *TreeNode {
	return &TreeNode{
		Value:      v,
	}

}