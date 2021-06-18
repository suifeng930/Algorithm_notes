package main

import (
	"Algorithm_notes/utils"
	"log"
	"math"
)

func main() {

	root := utils.NewNode(5)
	leftNode := utils.NewNode(1)
	rightNode := utils.NewNode(4)
	rleftNode1 := utils.NewNode(3)
	rrgihtNode2 := utils.NewNode(6)
	root.Left = leftNode
	root.Right = rightNode
	rightNode.Left = rleftNode1
	rightNode.Right = rrgihtNode2

	root.PrintTree()
	bst := isValidBST(root)
	log.Println(bst)
}

// https://leetcode-cn.com/problems/validate-binary-search-tree/
// 验证二叉搜索树
func isValidBST(root *utils.TreeNode) bool {

	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *utils.TreeNode, lower, upper int) bool {

	// 如果是空节点 返回true
	if root == nil {
		return true
	}

	// 判断当前节点是否在 上界 和下界内
	if root.Val <= lower {
		return false
	}
	if root.Val >= upper {
		return false
	}

	// 将当前节点的值替换为下界，继续检查右边的子节点
	if !helper(root.Right, root.Val, upper) {
		return false
	}
	// 将当前节点的值替换为上界，继续检查左边的子节点
	if !helper(root.Left, lower, root.Val) {
		return false
	}
	return true
}
