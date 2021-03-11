package main

import (
	"log"
)

func main() {

	stack := NewMinStack()

	stack.Push(4)
	stack.Push(9)
	stack.Push(7)
	stack.Push(3)
	stack.Push(8)
	stack.Push(5)
	stack.Push(2)
	log.Println(stack.GetMin())
	log.Println(stack.stack)
	log.Println(stack.minStack)
	stack.Pop()
	stack.Pop()
	stack.Pop()
	log.Println(stack.GetMin())
	log.Println(stack.stack)
	log.Println(stack.minStack)

	linkStack := NewLinkedList()
	linkStack.push(4)
	linkStack.push(9)
	linkStack.push(7)
	linkStack.push(3)
	linkStack.push(8)
	linkStack.push(5)
	linkStack.push(2)
	linkStack.Traverse()
	log.Println("linkstack min value :",linkStack.getMin())
	log.Println("linkstack top value :",linkStack.top())
	linkStack.pop()
	linkStack.pop()
	linkStack.pop()
	log.Println("linkstack min value :",linkStack.getMin())
	log.Println("linkstack top value :",linkStack.top())


}
