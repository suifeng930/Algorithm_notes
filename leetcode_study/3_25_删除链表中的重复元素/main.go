package main

func main() {

	listNode :=NewNode(1)
	listNode.Next=NewNode(2)
	listNode.Next.Next=NewNode(3)
	listNode.Next.Next.Next=NewNode(3)
	listNode.Next.Next.Next.Next=NewNode(4)
	listNode.Next.Next.Next.Next.Next=NewNode(4)
	listNode.Next.Next.Next.Next.Next.Next=NewNode(5)
	listNode.traverse()
	duplicates := deleteDuplicates(listNode)
	duplicates.traverse()

}

func NewNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}

}