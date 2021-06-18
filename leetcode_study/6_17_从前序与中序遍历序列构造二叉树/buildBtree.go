package main

import "Algorithm_notes/utils"

func main() {

	//前序遍历 preorder = [3,9,20,15,7]
	//中序遍历 inorder = [9,3,15,20,7]
	//    3
	//   / \
	//  9  20
	//    /  \
	//   15   7
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	treeNode := buildTree(preorder, inorder)

	treeNode.PrintTree()

}

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 从前序与中序遍序列构造二叉树
func buildTree(preOrder, inOrder []int) *utils.TreeNode {

	if len(inOrder) == 0 {
		return nil
	}
	root := &utils.TreeNode{Val: preOrder[0]}

	if len(inOrder) == 1 {
		return root
	}
	// 找到根节点索引位置
	//  前序遍历：    根-->左-->右
	//  中序遍历：    左-->根-->右
	rootIndex := findInOrderRootIndex(root.Val, inOrder)
	// 递归调用
	root.Left = buildTree(preOrder[1:rootIndex+1], inOrder[:rootIndex])
	root.Right = buildTree(preOrder[rootIndex+1:], inOrder[rootIndex+1:])
	return root
}

// 遍历inOrder 找到根节点所在的索引位置
func findInOrderRootIndex(rootVal int, inOrder []int) int {

	for key, val := range inOrder {
		if rootVal == val {
			return key
		}
	}
	return 0
}
