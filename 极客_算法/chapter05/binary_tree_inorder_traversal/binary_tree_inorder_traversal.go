package main


type TreeNode struct {

	Val int
	Left *TreeNode
	Right *TreeNode
}
// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
// 二叉树是中序遍历

func inorderTraversal(root *TreeNode) (res []int) {

	var inorder func(node *TreeNode)
	inorder= func(node *TreeNode) {
		if node==nil {
			return
		}
		// 先左子树
		inorder(node.Left)
		// 再根节点
		res =append(res,node.Val)
		// 最后右子树
		inorder(node.Right)
	}
	inorder(root)
	return res

}

// 中序递归
func inorderTraversalV2(root *TreeNode)  []int {

	res :=[]int{}
	inorder(root,&res)
	return res

}
// 中序递归
func inorder(root *TreeNode,res *[]int)  {
	//递归终止条件，节点为nil
	if root==nil {
		return
	}
	// 先遍历左子树
	inorder(root.Left,res)
	// 记录根节点
	*res=append(*res,root.Val)
	// 最后遍历右子树
	inorder(root.Right,res)
}


// 中序遍历 栈实现
func inorderTraversalV3(root *TreeNode)  []int {

	var res []int
	// 显示维护一个treeNode栈
	stack :=[]*TreeNode{}
	// 如果栈不为空 或者root节点不为nil
	for root!=nil || len(stack)>0{

		for root!=nil{
			// root 节点入栈，一直递归找root的left节点 并入栈，如找不到，则退出循环
			stack=append(stack,root)
			root=root.Left
		}
		// 节点出栈
		root=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		// 入数组中
		res =append(res,root.Val)
		// 遍历右子树
		root=root.Right
	}
	return res
}

/**
 算法伪码
 stack S;
 p=root;

 while(p || stack 不空）{
		while(p){
			p 入 stack;
			p=p的左子树；
		}
		p= stack.top 出栈;
		访问p;
		p=p的右子树；
}
 */


func NewNode(v int) *TreeNode {
	return &TreeNode{
		Val:      v,
	}

}