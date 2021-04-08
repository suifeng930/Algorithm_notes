package main

func main() {

	list :=&ListNode{
		Val:  1,
		Next: nil,
	}
	list1 :=& ListNode{
		Val:  3,
		Next: nil,
	}
	list2 :=& ListNode{
		Val:  4,
		Next: nil,
	}
	list3 :=& ListNode{
		Val:  1,
		Next: nil,
	}
	list4 :=& ListNode{
		Val:  2,
		Next: nil,
	}
	list5 :=& ListNode{
		Val:  4,
		Next: nil,
	}
	list.Next=list1
	list1.Next=list2

	list3.Next=list4
	list4.Next=list5

	printList(list)
	printList(list3)
	lists := mergeTwoLists(list, list3)
	printList(lists)
}
