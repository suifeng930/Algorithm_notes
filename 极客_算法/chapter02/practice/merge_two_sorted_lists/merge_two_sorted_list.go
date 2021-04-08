package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
// https://leetcode-cn.com/problems/merge-two-sorted-lists/
// 合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy :=&ListNode{Val: -1,Next: nil}
	prev :=dummy
	for l1 !=nil && l2!=nil {
		if l1.Val<=l2.Val {
			prev.Next=l1
			l1=l1.Next
		}else {
			prev.Next=l2
			l2=l2.Next
		}
		// prev 指针后移
		prev=prev.Next
	}
	// 合并l1和l2之后，最多只有一个还没有被合并完，我们直接将链表末尾指向未合并的链表即可
	if l1==nil {
		prev.Next=l2
	}else {
		prev.Next=l1
	}
	return dummy.Next
}

func printList(head *ListNode)  {
	curr :=head
	for curr!=nil {
		fmt.Printf("%d-->",curr.Val)
		next:=curr.Next
		curr=next
	}
	fmt.Println()
}