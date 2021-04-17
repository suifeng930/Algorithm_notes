package main

type Node struct {
	Val      int
	Children []*Node
}

// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
// 层序遍历 递归实现
func levelOrder(root *Node) [][]int {

	res := [][]int{}
	Order(0, root, &res)
	return res
}

// 递归子函数
func Order(level int, node *Node, res *[][]int) {

	// 初始化一个空数组
	val := []int{}
	// 递归退出条件
	if node == nil {
		return
	}
	// 追加层级
	if len(*res) <= level {
		*res = append(*res, val)
	}
	// 将数据追加到结果集合中
	(*res)[level] = append((*res)[level], node.Val)
	for _, child := range node.Children {
		Order(level+1, child, res)
	}
}

// N叉树的层序遍历
// 广度优先遍历   --->队列
func levelOrderV2(root *Node) [][]int  {

	// 判空处理
	if root==nil {
		return nil
	}
	// 初始化结果集，将root.Val 入结果集
	res :=[][]int{[]int{root.Val}}

	// 遍历root节点的子节点
	currentNodes :=root.Children

	for len(currentNodes)>0 {
		// 记录每一层子节点的个数
		k :=len(currentNodes)
		// 初始化每层的结果集
		levelArr :=make([]int,k)
		for i :=0;i<k;i++ {
			// 子节点结果入当前层的结果集中
			levelArr[i]=currentNodes[i].Val
			// 将当前节点的所有子节点入队列中
			currentNodes=append(currentNodes,currentNodes[i].Children...)
		}
		// 当前层结果集入 大集合
		res=append(res,levelArr)
		// 当前层节点出队列
		currentNodes=currentNodes[k:len(currentNodes)]
	}
	return res
}

func NewNode(v int) *Node {
	return &Node{
		Val: v,
	}

}
