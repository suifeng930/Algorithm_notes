package main

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

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
//  二叉树的前序遍历
//  递归实现
func preorderTraversal(root *TreeNode) []int {

	res := []int{}
	preOrder(root, &res)
	return res
}

//递归子问题
func preOrder(root *TreeNode, res *[]int) {
	// 递归终止条件
	if root == nil {
		return
	}
	// 先遍历根节点
	*res = append(*res, root.Val)
	// 再遍历左子树
	preOrder(root.Left, res)
	// 最后遍历右子树
	preOrder(root.Right, res)
}

// 栈实现
func preorderTraversalV2(root *TreeNode) []int {

	stack := []*TreeNode{}
	res := []int{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			// 先将根节点 和所有左孩子入栈，并加入到结果中，直到node 为空
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		// 然后，每弹出一个栈顶元素 ,就达到它的右孩子，在将这个节点当作  当前node 重新按照上面的步骤走一遍，直到栈为空。
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
		// 出栈
	}
	return res
}

// 栈实现
/**
  初始化栈，并将根节点入栈；
  当栈不为空时：
  	弹出栈顶元素node,并将值添加到结果集中
  	如果node的右子树非空，将右子树入栈
  	如果node的左子树非空，将左子树入栈
  由于栈是 先进后出的顺序，所以入栈时先将右子树入栈，这样使得前序遍历结果为 `根-->左-->右`的顺序
*/
func preorderTraversalV3(root *TreeNode) []int {

	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{}
	// 初始化栈，将根节点入栈
	stack = append(stack, root)
	node := &TreeNode{}
	for len(stack) > 0 {
		// 弹出栈顶元素node 并将值添加到结果集中
		node = stack[len(stack)-1]
		res = append(res, node.Val)
		stack = stack[:len(stack)-1]
		// 如果node的右子树非空，将右子树入栈
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		// 如果node的左子树非空，将左子树入栈
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}

func NewNode(v int) *TreeNode {
	return &TreeNode{
		Val: v,
	}

}
