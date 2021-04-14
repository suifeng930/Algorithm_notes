package main

import "log"

func main() {

	// input   [1, null, 2, 3]
	root := NewNode(1)
	rightNode := NewNode(2)
	leftNode := NewNode(3)
	root.Right = rightNode
	rightNode.Left = leftNode

	//v2 := preorderTraversal(root)
	//v2 := preorderTraversalV2(root)
	v2 := preorderTraversalV3(root)
	log.Println(v2)
	
}
