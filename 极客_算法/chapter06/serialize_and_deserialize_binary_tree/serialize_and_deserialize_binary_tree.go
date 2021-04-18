package main

import (
	"fmt"
	"strconv"
	"strings"
)

//https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/

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

type Codec struct {
	list []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return dfsSerial(root, "")
}

// 前序遍历 根-->左-->右
func dfsSerial(root *TreeNode, str string) string {
	if root == nil {  // 遍历到null 节点
		str += "null,"
	} else {
		str += strconv.Itoa(root.Val) + "," // 拼接根节点val
		str = dfsSerial(root.Left, str) // 左子树序列化结果
		str = dfsSerial(root.Right, str) // 右子树序列化结果
	}
	return str
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	queue := strings.Split(data, ",")
	return buildTree(&queue)
}

// 前序遍历 根-->左-->右
func buildTree(queue *[]string) *TreeNode {

	rootVal := (*queue)[0]
	*queue = (*queue)[1:]
	if rootVal == "null" {
		return nil
	}
	val, _ := strconv.Atoi(rootVal) // 当前节点不为null
	root := &TreeNode{Val: val}  //初始化一个root节点
	root.Left = buildTree(queue) // 递归构建root节点的左子树
	root.Right = buildTree(queue)// 递归构建root节点的右子树
	return root
}




// Serializes a tree to a single string.
func (this *Codec) serializeV2(root *TreeNode) string {

	queue :=[]*TreeNode{}
	queue=append(queue,root)
	res :=[]string{}
	for len(queue)>0 {
		node :=queue[0]  // 出队列
		queue=queue[1:]
		if node!=nil {  // 节点不为空，
			res =append(res,strconv.Itoa(node.Val)) //将根节点加入结果集中
			queue=append(queue,node.Left) //左子节点入队
			queue=append(queue,node.Right) //右子节点入队
		}else {
			res=append(res,"null") //如果是nil节点 用null代替
		}
	}
	return strings.Join(res,",")
}



// Deserializes your encoded data to tree.
func (this *Codec) deserializeV2(data string) *TreeNode {

	if data=="null" {
		return nil
	}
	list :=strings.Split(data,",")// 序列化字符串split成数组
	val, _ := strconv.Atoi(list[0])
	root :=&TreeNode{Val: val}   // 构建根节点
	queue :=[]*TreeNode{}
	queue=append(queue,root) // 根节点入队列
	cursor :=1    // data指向下一个节点
	for cursor<len(list) {
		node :=queue[0]
		queue=queue[1:]  // 出队列
		leftVal :=list[cursor]
		rightVal:=list[cursor+1]
		if leftVal!="null" { // 如果左子节点val 不为空 入队列
			Val, _ := strconv.Atoi(leftVal)
			leftNode:=&TreeNode{Val: Val}
			node.Left=leftNode
			queue=append(queue,leftNode)
		}

		if rightVal!="null" { // 如果右子节点val 不为空入队列
			Val, _ := strconv.Atoi(rightVal)
			rightNode :=&TreeNode{Val: Val}
			node.Right=rightNode
			queue=append(queue,rightNode)
		}
		cursor+=2
	}
	return root
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

 // 前序遍历 根-->左-->右
func PrintTree(root *TreeNode) {

	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {

		for cur != nil {
			fmt.Printf("%v-->", cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = cur.Right
	}
	fmt.Println()
}

func NewNode(v int) *TreeNode {
	return &TreeNode{
		Val: v,
	}

}
