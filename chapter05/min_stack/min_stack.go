package main

import (
	"fmt"
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


	log.Println("测试一个字符串读前序遍历")
	str :="9,3,4,#,#,1,#,#,2,#,6,#,#"
	serialization := isValidSerialization(str)
	log.Println(serialization)
	v2 := isValidSerializationV2(str)
	log.Println(v2)
	//str1 :="1,#"
	str2 :="9,#,92,#,#"
	v3 := isValidSerializationV3(str2)
	log.Println(v3)

	var a uint =1
	var b uint =2
	fmt.Printf("%d\n", a-b)
	//maxUint64 := math.MaxUint64
	//fmt.Printf("%d\n", maxUint64)
}
