package main

import "log"

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

func main() {

	root := NewNode(1)
	leftNode := NewNode(2)
	rightNode := NewNode(3)
	rrgihtNode1:=NewNode(4)
	rrgihtNode2:=NewNode(5)
	root.Left=leftNode
	root.Right=rightNode
	rightNode.Left=rrgihtNode1
	rightNode.Right=rrgihtNode2
	PrintTree(root)

	ser := Constructor()
	deser:= Constructor()
	data := ser.serialize(root)
	log.Println(data)
	ans := deser.deserialize(data)
	PrintTree(ans)
	log.Println(ans)
}
