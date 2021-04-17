package main

/**
*      4
*    /   \
*   2     7
*  / \   / \
* 1   3 6   9
*
*/

func main() {
	root := NewNode(4)
	leftNode := NewNode(2)
	rightNode := NewNode(7)
	lleftNode1 :=NewNode(1)
	lleftNode2 :=NewNode(3)
	rrightNode1 :=NewNode(6)
	rrightNode2 :=NewNode(9)
	root.Left=leftNode
	root.Right=rightNode
	leftNode.Left=lleftNode1
	leftNode.Right=lleftNode2
	rightNode.Left=rrightNode1
	rightNode.Right=rrightNode2

	// 前序遍历打印翻转前的树
	PrintTreePre(root)
	PrintTree(root)
	v2 := invertTree(root)
	// 前序遍历打印翻转后的树
	PrintTreePre(v2)
}
