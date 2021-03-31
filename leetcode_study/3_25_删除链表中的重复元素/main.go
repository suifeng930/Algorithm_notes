package main

func main() {

	listNode := NewNode(1)
	listNode.Next = NewNode(2)
	listNode.Next.Next = NewNode(3)
	listNode.Next.Next.Next = NewNode(3)
	listNode.Next.Next.Next.Next = NewNode(4)
	listNode.Next.Next.Next.Next.Next = NewNode(4)
	listNode.Next.Next.Next.Next.Next.Next = NewNode(5)
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

func maxArea(height []int) int {
	maxVal := 0
	for i, j := 0, len(height)-1; i < j; {
		minHeight := 0
		if compare2Min(height[i], height[j]) {
			minHeight = height[i]
			i++
		} else {
			minHeight = height[j]
			j--
		}
		area := (j - i + 1) * minHeight
		maxVal = compare2Max(maxVal, area)
	}
	return maxVal
}

func compare2Min(x, y int) bool {
	if x < y {
		return true
	}
	return false
}

func compare2Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

