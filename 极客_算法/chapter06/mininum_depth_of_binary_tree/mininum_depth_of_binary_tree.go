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

//  https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

// 二叉树的最小深度
//  深度优先遍历
func minDepth(root *TreeNode) int {

	// 如果根节点为nil  返回0
	if root==nil{
		return 0
	}
	// 如果左子树或右子树为nil 直接返回1
	if root.Left==nil && root.Right==nil {
		return 1
	}
	minD :=math.MaxInt64
	// 如果左子树不为nil ,递归寻找最小值
	if root.Left!=nil{
		minD=min(minDepth(root.Left),minD)
	}
	// 如果右子树不为nil ,递归寻找最小值
	if root.Right!=nil {
		minD=min(minDepth(root.Right),minD)
	}
	return minD+1
}

func min(a int, b int) int {
	if a<b {
		return a
	}
	return b
}


// 二叉树的最小深度
//  广度优先遍历(层序遍历)
func minDepthV2(root *TreeNode) int {

	if root==nil {
		return 0
	}

	queue :=[]*TreeNode{}
	queue=append(queue,root)
	minN :=0
	for len(queue)>0  {
		minN++
		newQueue :=make([]*TreeNode,0)
		for _, node := range queue{
			// 节点的左、右节点都为nil 说明已经找到子节点了。
			if node.Left==nil && node.Right==nil {
				return minN
			}
			if node.Left!=nil {
				newQueue=append(newQueue,node.Left)
			}
			if node.Right!=nil {
				newQueue=append(newQueue,node.Right)
			}
		}
		//遍历每一层的节点，每次都会更新每一层的队列
		queue=newQueue
	}
	return minN
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