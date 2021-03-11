package main

import "log"

type LinkedList struct {
	headNode *Node
}

type Node struct {
	value    int
	minValue int
	next     *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil}
}

func(l *LinkedList) IsEmpty() bool {
	return l.headNode==nil
}

func (l *LinkedList) push(x int) {
	if l.IsEmpty() {
		newNode :=&Node{
			value:    x,
			minValue: x,
			next:     nil,
		}
		l.headNode=newNode
	}else {
		minValue := CompareValue(x, l.headNode.minValue)
		newNode :=&Node{
			value:    x,
			minValue: minValue,
			next:     l.headNode,
		}
		l.headNode=newNode
	}
}

func (l *LinkedList) pop() {
	if l.IsEmpty() {
		return
	}
	popValue :=l.headNode.value
	log.Println("pop value :",popValue)
	l.headNode=l.headNode.next
}

func (l *LinkedList) top() int {
	return l.headNode.value
}

func (l *LinkedList) getMin() int {
	return l.headNode.minValue
}

func (l *LinkedList) Traverse() {
	if l.IsEmpty() {
		log.Println("栈已空")
		return
	}
	currentNode :=l.headNode
	for currentNode!=nil {
		log.Printf("%v\t",currentNode.value)
		currentNode=currentNode.next
	}
}


// 比较元素大小
func CompareValue(x, y int) int {
	if x < y {
		return x
	}
	return y
}
