package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	// 如果链表为空，返回
	if head == nil {
		return nil
	}
	//初始化一个哑巴节点，并将哑巴节点指向头节点
	dummy := &ListNode{0, head}
	// 将cur指针，指向哑巴节点
	cur := dummy
	// 遍历链表 cur.next
	for cur.Next != nil && cur.Next.Next != nil {

		//如果 cur.next.val 等于cur.next.next.val 说明数据相同
		if cur.Next.Val == cur.Next.Next.Val {
			// 将cur.next.val 记录下来
			x := cur.Next.Val
			// 不断遍历 找到和val值相等的节点，并删除
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			// 说明数字不相等，cur指针往下移动
			cur = cur.Next
		}
	}
	return dummy.Next
}

func (l *ListNode) traverse() {

	if l == nil {
		return
	}
	currentNode := l
	for currentNode != nil {
		fmt.Printf("%v->\t", currentNode.Val)
		currentNode = currentNode.Next
	}
	fmt.Println()
}
