package main

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	Stack []*NestedInteger //用于存储数据的栈 slice
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	stack :=[]*NestedInteger{}  //初始化一个slice
	for i :=len(nestedList)-1;i>=0;i--{ // 从后向前遍历给定的数列，因为栈是先进后出，如果在入栈的时候从后向前入栈，那么最后一个元素就在栈底了，相应的栈顶其实就是存储的第一个元素
		stack=append(stack,nestedList[i])//
	}
	return &NestedIterator{stack}
}

func (this *NestedIterator) Next() int {
	this.stackTop2Integer() //先将栈顶元素转换成integer
	top :=this.Stack[len(this.Stack)-1] //获取栈顶元素
	this.Stack=this.Stack[:len(this.Stack)-1] //出栈
	return top.GetInteger() //返回栈顶元素
}

func (this *NestedIterator) HasNext() bool {
	this.stackTop2Integer() //将栈顶元素转换成integer
	return len(this.Stack)>0 //判断栈是否为空，
}
// 将栈顶的元素转化成为integer 如果栈顶元素不是integer，则说明栈顶是一个list,此时需要遍历这个list，并将list中的元素依次逆序入栈
func(this *NestedIterator) stackTop2Integer(){

	for len(this.Stack)>0{ //边界条件，如果栈不为空，则一直循环
		top :=this.Stack[len(this.Stack)-1] // 获取栈顶元素
		if top.IsInteger(){ //判断栈顶元素是否为integer 如果是，则返回
			return
		}
		this.Stack=this.Stack[:len(this.Stack)-1] // 当前栈顶元素是列表,弹出当前栈顶元素
		list :=top.GetList() //将该列表转换成一个list，
		for i :=len(list)-1;i>=0;i--{ //从后向前遍历，依次入栈
			this.Stack=append(this.Stack,list[i])
		}
	}
}