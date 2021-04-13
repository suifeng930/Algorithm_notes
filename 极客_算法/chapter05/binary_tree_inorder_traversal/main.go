package main

import "log"

func main() {

	root := NewNode(1)
	rightNode := NewNode(2)
	leftNode := NewNode(3)
	root.Right=rightNode
	rightNode.Left=leftNode

	//v2 := inorderTraversalV2(root)
	v2 := inorderTraversalV3(root)
	log.Println(v2)

}
