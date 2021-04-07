package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


func reverseKGroup(head *ListNode, k int) *ListNode {

	dummy :=&ListNode{Val:0,Next: head}
	pre :=dummy

	for head!=nil {
		// 定义子链表的尾节点
		tail :=pre
		for i :=0;i<k;i++{
			tail=tail.Next
			if tail==nil {
				return dummy.Next
			}
		}
		// 下一个子链表的头节点
		temp :=tail.Next
		// 调用子链表翻转， 返回翻转之后的，头 、尾指针
		head,tail =myReverse(head,tail)
		pre.Next=head
		tail.Next=temp
		pre=tail
		head=tail.Next
	}
	return dummy.Next
}

//局部翻转操作
func myReverse(head, tail *ListNode) (*ListNode,*ListNode) {

	// 下一个子链表的头节点
	prev :=tail.Next
	//当前子链表的头节点
	p:=head
	// 完成k次局部翻转操作
	for prev!=tail {
		// temp  指针 指向头节点的下一个节点，
		nex :=p.Next
		// 将头节点的下一个指针，指向下一个子链表的头指针
		p.Next=prev
		// 头节点向后移动
		prev=p
		// 尾节点向前移
		p=nex
	}
	return tail,head
}

func printList(head *ListNode)  {

	curr :=head
	for curr!=nil {
		fmt.Printf("%d-->",curr.Val)
		next :=curr.Next
		curr=next
	}
	fmt.Println()
}