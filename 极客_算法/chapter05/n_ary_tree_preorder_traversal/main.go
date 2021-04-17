package main

import "log"

func main() {

	root := NewNode(1)

	firstChild := NewNode(3)
	secondChild := NewNode(2)
	threeChild := NewNode(4)
	Child1 := NewNode(5)
	Child2 := NewNode(6)
	root.Children=append(root.Children,firstChild,secondChild,threeChild)
	firstChild.Children=append(firstChild.Children,Child1,Child2)

	vals := preorder(root)
	log.Println(vals)
}
