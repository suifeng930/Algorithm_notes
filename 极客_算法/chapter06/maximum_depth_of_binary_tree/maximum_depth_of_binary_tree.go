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


// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

// 二叉树的最大深度
// 递归深度优先搜索
func maxDepth(root *TreeNode) int {

	// 递归终止条件
	if root==nil {
		return 0
	}
	// 递归条件
	return max(maxDepth(root.Left),maxDepth(root.Right))+1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 二叉树的最大深度  广度优先遍历(层序遍历)
func maxDepthV2(root *TreeNode) int {
	if root==nil {
		return 0
	}
	queue :=[]*TreeNode{}
	queue=append(queue,root)
	level :=0
	for len(queue)>0 {
		size := len(queue)
		for size>0 {
			node :=queue[0] // queue.peek
			queue=queue[1:] // queue.pop
			if node.Left!=nil {
				queue=append(queue,node.Left)
			}
			if node.Right!=nil {
				queue=append(queue,node.Right)
			}
			size--
		}
		level++
	}
	return level
}


func PrintTree(root *TreeNode)  {
	stack :=[]*TreeNode{}

	for root!=nil || len(stack)>0 {

		for root!=nil {
			stack=append(stack,root)
			root=root.Left
		}
		root=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		fmt.Printf("%d-->",root.Val)
		root=root.Right
	}
	fmt.Println()
}

func NewNode(v int) *TreeNode {
	return &TreeNode{
		Val:      v,
	}
}