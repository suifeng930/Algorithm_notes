package main

import "fmt"

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

//  https://leetcode-cn.com/problems/invert-binary-tree/

// 翻转二叉树 递归求解
func invertTree(root *TreeNode) *TreeNode {

	// 递归终止条件
	if root==nil {
		return nil
	}
	// 递归逻辑
	//  翻转 根节点的左子树
	leftTree :=invertTree(root.Left)
	//  翻转 根节点的右子树
	rightTree :=invertTree(root.Right)
	//  将左子树和右子树翻转
	root.Left=rightTree
	root.Right=leftTree
	//  返回结果集
	return root
}

// 二叉树的翻转
func invertTreeV2(root *TreeNode) *TreeNode {

	// 递归终止条件
	if root==nil {
		return nil
	}
	// 翻转左子树
	invertTree(root.Left)
	// 翻转右子树
	invertTree(root.Right)
	// 交换左、右子节点
	root.Left,root.Right=root.Right,root.Left

	return root


}



// 二叉树的中序遍历
func PrintTree(root *TreeNode) {

	stack :=[]*TreeNode{}
	cur :=root
	for cur!=nil|| len(stack)>0 {
		for cur!=nil {
			stack=append(stack,cur)
			cur=cur.Left
		}
		cur=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		fmt.Printf("%d-->",cur.Val)
		cur=cur.Right
	}
	fmt.Println()
}

func PrintTreePre(root *TreeNode)  {
	stack :=[]*TreeNode{}
	cur :=root
	for cur!=nil ||len(stack)>0 {

		for cur!=nil {
			fmt.Printf("%d-->",cur.Val)
			stack=append(stack,cur)
			cur=cur.Left
		}
		cur=stack[len(stack)-1].Right
		stack=stack[:len(stack)-1]
	}
	fmt.Println()
}



func NewNode(v int) *TreeNode {
	return &TreeNode{
		Val:      v,
	}

}