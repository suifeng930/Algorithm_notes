package main


func main() {

	list :=ListNode{
		Val:  1,
		Next: nil,
	}
	list1 :=& ListNode{
		Val:  2,
		Next: nil,
	}
	list2 :=& ListNode{
		Val:  3,
		Next: nil,
	}
	list3 :=& ListNode{
		Val:  4,
		Next: nil,
	}
	list.Next=list1
	list1.Next=list2
	list2.Next=list3

	printList(&list)
	listNode := reverseList(&list)
	printList(listNode)

}
