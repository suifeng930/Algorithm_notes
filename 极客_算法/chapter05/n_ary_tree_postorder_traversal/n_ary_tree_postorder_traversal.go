package main

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

// https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/
// n 叉树的后序遍历

func postorder(root *Node) []int {

	res := []int{}
	order(root, &res)
	return res
}

func order(root *Node, res *[]int) {

	// 递归终结条件
	if root == nil {
		return
	}
	// 当层循环 主体
	for _, child := range root.Children {
		order(child, res)
	}
	// 将当前val 追加到 res集合中
	*res = append(*res, root.Val)

}

// 使用栈 解决
func postOrder(root *Node) []int {
	res := []int{}

	if root == nil {
		return res
	}
	stack := []*Node{}
	//根节点入栈
	stack = append(stack, root)
	for len(stack) > 0 {
		// 根节点出栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		//将根节点的子节点全部入栈
		stack = append(stack, root.Children...)
	}
	// 反转结果集
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

func Print(root *Node) {

	node := root
	for node != nil {
		fmt.Printf("%d-->\t", node.Val)

	}
}

func NewNode(v int) *Node {
	return &Node{
		Val: v,
	}

}
