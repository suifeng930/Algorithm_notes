package main

import (
	"Algorithm_notes/utils"
	"log"
)

func main() {
	root := utils.NewNode(3)
	leftNode := utils.NewNode(9)
	rightNode := utils.NewNode(20)
	rrgihtNode1 := utils.NewNode(15)
	rrgihtNode2 := utils.NewNode(7)
	root.Left = leftNode
	root.Right = rightNode
	rightNode.Left = rrgihtNode1
	rightNode.Right = rrgihtNode2

	root.PrintTree()
	depth := maxDepth(root)
	log.Println(depth)

}

//方法二：层序遍历（BFS）
//树的层序遍历 / 广度优先搜索往往利用 队列 实现。
//关键点： 每遍历一层，则计数器 +1+1 ，直到遍历完成，则可得到树的深度。
//算法解析：
//特例处理： 当 root​ 为空，直接返回 深度 00 。
//初始化： 队列 queue （加入根节点 root ），计数器 res = 0。
//循环遍历： 当 queue 为空时跳出。
//初始化一个空列表 tmp ，用于临时存储下一层节点；
//遍历队列： 遍历 queue 中的各节点 node ，并将其左子节点和右子节点加入 tmp；
//更新队列： 执行 queue = tmp ，将下一层节点赋值给 queue；
//统计层数： 执行 res += 1 ，代表层数加 11；
//返回值： 返回 res 即可。
//

//https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/submissions/
// 层序遍历。广度优先遍历
func maxDepth(root *utils.TreeNode) int {

	if root == nil {
		return 0
	}
	depth := 0
	queue := []*utils.TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 { // queue为空 跳出循环
		cnt := len(queue) // 记录每一层的节点树
		for cnt > 0 {
			node := queue[0]  // queue.peek
			queue = queue[1:] // queue.pop
			// 如果左右节点不为nil ,入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			cnt--
		}
		depth++
	}
	return depth
}
