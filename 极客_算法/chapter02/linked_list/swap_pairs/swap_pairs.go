package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

//  https://leetcode-cn.com/problems/swap-nodes-in-pairs/
//  两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {

	// 创建一个哑巴节点，next指针指向 head
	dummyHead :=&ListNode{Val:  0, Next: head}
	// 初始化temp节点，指向dummyHead节点，每次交换temp.next temp.next.next
	temp :=dummyHead

	//边界条件， 1.temp后面没有节点，即temp.next ==nil
	//         2.temp后面只有一个节点，即temp.next.next ==nil
	for temp.Next !=nil && temp.Next.Next!=nil {
		// node1 即为 temp.next
		node1 :=temp.Next
		// node2 即为 temp.next.next eg : 0->1->2->3->4
		node2:=temp.Next.Next
		// 1. temp.next指针指向node2  eg: 0->2
		temp.Next=node2
		// 2. node1.next指针指向node2.next eg: 1->3
		node1.Next=node2.Next
		// 3. node2.next指针指向node1 eg: 2->1
		//  完成局部反转之后： eg : 0->2->1->3->4
		node2.Next=node1
		// 4. 将temp指针前移至node1  temp指针即指向节点1 eg: 1->3->4  继续迭代
		temp=node1
	}
	// 返回的是dummyHead.next指针，排序后的头节点
	return dummyHead.Next
}

// 递归求解
// 1. 返回值： 交换完成的子链表
// 2. 调用单元：设需要交换的两个点为head 和next，head连接后面交换完成的子链表，next连接head,完成交换
// 3. 终止条件：head为空指针或者next为空指针，也就是当前无节点或者只有一个节点，无法进行交换
func swapPairsV2(head *ListNode) *ListNode {

	if head==nil || head.Next==nil {
		return head
	}
	newHead :=head.Next
	head.Next=swapPairs(newHead.Next)
	newHead.Next=head
	return newHead
}


func printList(head *ListNode)  {
	curr :=head
	for curr!=nil {
		fmt.Printf("%d-->\t",curr.Val)
		next :=curr.Next
		curr=next
	}
	fmt.Println()
}


