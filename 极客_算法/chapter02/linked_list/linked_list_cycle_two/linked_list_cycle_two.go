package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {

	curr := head
	for curr != nil {
		fmt.Printf("%d-->\t", curr.Val)
		next := curr.Next
		curr = next
	}
}
//  https://leetcode-cn.com/problems/linked-list-cycle-ii/
func detectCycle(head *ListNode) *ListNode {

	// 定义快慢指针
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 快慢指针相遇
		if fast == slow {
			//初始化一个ptr指针，从头节点向后移动，
			ptr := head
			//直到slow指针和ptr指针相遇，相遇点即为入环点
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			return ptr
		}
	}
	return nil
}
