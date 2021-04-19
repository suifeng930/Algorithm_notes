package main

import "log"

func main() {

	/**
	#                      3
	#                    /  \
	#                   5    1
	#                  / \  / \
	#                 6  2  0  8
	#                   / \
	#                  7   4
	 */
	// 层序遍历的字符串
	rootStr :="3,5,1,6,2,0,8,null,null,7,4"
	root := Deserialize(rootStr)
	PrintTree(root)
	p :=NewNode(5)
	q :=NewNode(4)

	//ancestorNode := lowestCommonAncestor(root, p, q)
	ancestorNode := lowestCommonAncestorV2(root, p, q)

	log.Println(ancestorNode.Val)
}
