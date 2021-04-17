package main

type Node struct {
	Val int
	Children []*Node
}

// https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
// N叉树的前序遍历
func preorder(root *Node) []int  {
	res :=[]int{}
	preOrder(root,&res)
	return res

}

func preOrder(root *Node, res *[]int) {

	// 循环退出条件
	if root==nil {
		return
	}
	// 将root.val 入结果集
	*res=append(*res,root.Val)
	// 循环遍历子节点
	for _, child := range root.Children {
		preOrder(child,res)
	}
}

// 使用栈实现
func preorderV2(root *Node) []int  {

	stack :=[]*Node{}
	res :=[]int{}
	if root==nil {
		return res
	}
	stack=append(stack,root)
	for len(stack)>0 {
		root =stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		res=append(res,root.Val)
		// 注意这里遍历子节点的入栈顺序是 v3-->v2-->v1
		// 出栈顺序也就相应地变成了 v1-->v2-->v3 了
		for i := len(root.Children)-1;i>=0;i-- {
			stack=append(stack,root.Children[i])
		}
	}
	return res
}


func NewNode(v int) *Node {
	return &Node{
		Val: v,
	}

}
