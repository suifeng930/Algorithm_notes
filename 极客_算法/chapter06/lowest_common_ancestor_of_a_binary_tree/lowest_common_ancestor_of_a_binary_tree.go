package main

import (
	"fmt"
	"strconv"
	"strings"
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
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree
// 二叉树的最近公共祖先
// 递归求解
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	// 1. 如果根节点为nil 直接返回
	if root == nil {
		return nil
	}
	// 如果root的值 等于 p或者 q,返回root节点
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	// 递归root 左、右子树节点
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 左、右子树结果集不为null,则说明root就是最近公共祖先
	if left != nil && right != nil {
		return root
	}
	// 如果 left为nil 或者right不为nil则 最近公共祖先只可能出现在右子树中
	if left == nil {
		return right
	}
	// 如果left不为nil,且right为nil，则最近公共祖先只可能出现在左子树中。
	return left
}

// 二叉树的最近公共祖先
// hashtable
func lowestCommonAncestorV2(root, p, q *TreeNode) *TreeNode {

	parent :=map[int]*TreeNode{}
	visited :=map[int]bool{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node==nil {
			return
		}
		if node.Left!=nil {
			parent[node.Left.Val]=node
			dfs(node.Left)
		}
		if node.Right!=nil {
			parent[node.Right.Val]=node
			dfs(node.Right)
		}
	}
	dfs(root)
	for p!=nil {
		visited[p.Val]=true
		p=parent[p.Val]
	}
	for q!=nil {
		if visited[q.Val] {
			return q
		}
		q=parent[q.Val]
	}
	return nil
}

func PrintTree(root *TreeNode) {

	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d-->", root.Val)
		root = root.Right
	}
	fmt.Println()
}

// Deserializes your encoded data to tree.
// 层序遍历 构造一棵树
func Deserialize(data string) *TreeNode {
	if data == "null" {
		return nil
	}
	list := strings.Split(data, ",")
	val, _ := strconv.Atoi(list[0])
	root := &TreeNode{Val: val}
	queue := []*TreeNode{}
	queue = append(queue, root)
	cursor := 1
	for cursor < len(list) {
		node := queue[0]
		queue = queue[1:]
		leftVal := list[cursor]
		rightVal := list[cursor+1]
		if leftVal != "null" {
			Val, _ := strconv.Atoi(leftVal)
			leftNode := &TreeNode{Val: Val}
			node.Left = leftNode
			queue = append(queue, leftNode)
		}
		if rightVal != "null" {
			Val, _ := strconv.Atoi(rightVal)
			rightNode := &TreeNode{Val: Val}
			node.Right = rightNode
			queue = append(queue, rightNode)
		}
		cursor += 2
	}
	return root
}
func NewNode(v int) *TreeNode {
	return &TreeNode{
		Val: v,
	}
}
