package main

import "log"

func main() {


	root := NewNode(3)
	leftNode := NewNode(9)
	rightNode := NewNode(20)
	rrgihtNode1:=NewNode(15)
	rrgihtNode2:=NewNode(7)
	root.Left=leftNode
	root.Right=rightNode
	rightNode.Left=rrgihtNode1
	rightNode.Right=rrgihtNode2

	PrintTree(root)
	//depth := minDepth(root)
	depth := minDepthV2(root)
	log.Println(depth)
}
