package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
// 反转链表
func reverseList(head *ListNode) *ListNode {

	// 0.prev 节点指向nil
	var prev *ListNode
	// 0.curr 指向链表的头节点，遍历链表
	curr := head
	for curr != nil {
		// 1.构造一个temp变量指向 curr.next指针
		next := curr.Next
		// 2.将curr 的next指针指向prev节点 (指针方向改变)
		curr.Next = prev
		// 3.将curr赋值给prev, 实现prev指针的向前移动
		prev = curr
		// 3.同理，将curr.next节点赋值给curr,实现curr指针向前移动
		curr = next
	}
	return prev
}


// 简便写法，需要理解
func reverseListV2(head *ListNode) *ListNode {

	// 0.prev 节点指向nil
	var prev *ListNode
	for head != nil {
		// 1.每次的操作都是：
		//将prev指针，赋值给head.Next ;
		//将head指针赋值给prev实现prev指针前移；
		//并将head.Next指针赋值给head,实现head指针前移
		// 此为上一个方法的简化版本
		head.Next,prev,head=prev,head,head.Next
	}
	return prev
}


func printList(head *ListNode) {

	curr := head
	for curr != nil {
		fmt.Printf("%d --->\t", curr.Val)
		next := curr.Next
		curr = next
	}
	fmt.Println()
}

type ListNode struct {
	Val  int
	Next *ListNode
}
