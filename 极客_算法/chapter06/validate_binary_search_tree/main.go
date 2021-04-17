package main

import "log"

func main() {

	root := NewNode(2)
	leftNode := NewNode(1)
	rightNode := NewNode(3)
	root.Left=leftNode
	root.Right=rightNode

	//v2 := isValidBST(root)
	v2 := isValidBSTV2(root)
	PrintTree(root)
	log.Println(v2)
}
