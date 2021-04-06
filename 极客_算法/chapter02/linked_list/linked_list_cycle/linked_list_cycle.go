package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

// https://leetcode-cn.com/problems/linked-list-cycle/
// 判断链表有环

func hasCycle(head *ListNode) bool {

	if head ==nil || head.Next==nil {
		return false
	}

	slow ,fast :=head,head.Next
	for fast!=slow {
		if fast==nil||fast.Next==nil {
			return false
		}
		slow=slow.Next
		fast=fast.Next.Next
	}
	return true
}

func printList(head *ListNode) {

	curr := head
	for curr != nil {
		fmt.Printf("%d-->\t", curr.Val)
		next := curr.Next
		curr = next
	}
	fmt.Println()
}