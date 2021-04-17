package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/validate-binary-search-tree/submissions/

// 验证二叉搜索树
func isValidBST(root *TreeNode) bool {

	return helper(root,math.MinInt64,math.MaxInt64)
}

func helper(root *TreeNode, lower int, upper int) bool {

	// 如果是空节点 返回true
	if root==nil {
		return true
	}
	// 判断节点是否在上界和下界内
	if root.Val<=lower {
		return false
	}
	if root.Val>=upper {
		return false
	}

	// 将当前节点的值替换为下界，继续检查右边的子节点
	if !helper(root.Right,root.Val,upper) {
		return false
	}
	// 将当前节点的值替换为上界，继续检查左边的子节点
	if !helper(root.Left,lower,root.Val) {
		return false
	}
	return true
}

// 验证二叉搜索树 中序遍历，值判断
func isValidBSTV2(root *TreeNode) bool {

	stack :=[]*TreeNode{}
	inOrder :=math.MinInt64

	for root!=nil || len(stack)>0 {

		for root!=nil {
			stack=append(stack,root)
			root=root.Left
		}
		root=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		// 如果当前节点值小于上一个节点，则返回false
		if root.Val<=inOrder {
			return false
		}
		inOrder=root.Val
		root=root.Right
	}
	return true
}




func PrintTree(root *TreeNode)  {

	stack :=[]*TreeNode{}
	cur :=root
	for cur!=nil ||len(stack)>0 {

		for cur!=nil {
			stack=append(stack,cur)
			cur=cur.Left
		}
		cur=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		fmt.Printf("%v-->",cur.Val)
		cur=cur.Right
	}
	fmt.Println()
}

func NewNode(v int) *TreeNode {
	return &TreeNode{
		Val:      v,
	}

}