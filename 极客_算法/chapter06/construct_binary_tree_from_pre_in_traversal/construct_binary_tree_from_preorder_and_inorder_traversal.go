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

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// 从前序遍历和后序遍历 获取一个树
// 递归实现：
func buildTree(preorder []int, inorder []int) *TreeNode {
	return helper(0,0, len(inorder)-1,preorder,inorder)

}

func helper(preStart int, inStart int, inEnd int, preorder []int, inorder []int) *TreeNode {

	// 递归终止条件， preStart 越界和inStart 越界
	if preStart> len(preorder)-1 || inStart>inEnd{
		return nil
	}
	// 当前层处理逻辑
	//  找到根节点， pre 前序遍历 根-->左-->右
	//             in  中序遍历 左-->根-->右
	root :=&TreeNode{Val: preorder[preStart]}
	inIndex :=0
	// 找到根节点的在inOrder中的索引位置
	for i:=inStart;i<=inEnd;i++ {
		if inorder[i]==root.Val {
			inIndex=i
		}
	}
	//递归进入下一层
	// 根节点的左子树为， inOrder[0,inIndex-1]的 区域
	root.Left=helper(preStart+1,inStart,inIndex-1,preorder,inorder)
	// 根节点的右子树为： inOrder[inIndex+1, end] 区域
	root.Right=helper(preStart+inIndex-inStart+1,inIndex+1,inEnd,preorder,inorder)
	return root
}

//  递归遍历
func buildTreeV2(preorder []int, inorder []int) *TreeNode {

	if len(inorder)==0 {
		return nil
	}
	root :=&TreeNode{Val: preorder[0]}

	if len(inorder)==1 {
		return root
	}
	// 找到inOrder的索引位置
	inIndex :=findInOrderRootIndex(root.Val,inorder)
	// 递归调用
	root.Left=buildTreeV2(preorder[1:inIndex+1],inorder[:inIndex])
	root.Right=buildTreeV2(preorder[inIndex+1:],inorder[inIndex+1:])
	return root
}
// 遍历inOrder 找到根节点所在的索引位置
func findInOrderRootIndex(rootVal int, inOrder []int) int {
	for i, val := range inOrder {
		if rootVal==val {
			return i
		}
	}
	return 0
}

func buildTreeV3(preorder []int, inorder []int) *TreeNode {

	if len(preorder)==0 {
		return nil
	}
	root :=&TreeNode{Val: preorder[0]}
	stack :=[]*TreeNode{}
	stack=append(stack,root)

	inOrderIndex :=0

	for i :=1;i<len(preorder);i++{
		preOrderVal := preorder[i]
		node :=stack[len(stack)-1]

		if node.Val!=inorder[inOrderIndex] {
			node.Left=&TreeNode{Val: preOrderVal}
			stack=append(stack,node.Left)
		}else {
			for len(stack)!=0 && stack[len(stack)-1].Val==inorder[inOrderIndex] {
				node=stack[len(stack)-1]
				stack=stack[:len(stack)-1]
				inOrderIndex++
			}
			node.Right=&TreeNode{Val: preOrderVal}
			stack=append(stack,node.Right)
		}
	}
	return root
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
