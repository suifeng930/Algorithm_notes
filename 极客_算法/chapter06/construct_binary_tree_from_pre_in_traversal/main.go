package main

func main() {
	//
	//前序遍历 preorder = [3,9,20,15,7]
	 //中序遍历 inorder = [9,3,15,20,7]
	preorder :=[]int{3,9,20,15,7}
	inorder :=[]int{9,3,15,20,7}

	tree := buildTree(preorder, inorder)
	PrintTree(tree)
	v2 := buildTreeV2(preorder, inorder)
	PrintTree(v2)
	v3 := buildTreeV3(preorder, inorder)
	PrintTree(v3)

}
