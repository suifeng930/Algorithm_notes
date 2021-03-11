package main

import "math"

type MinStack struct {
	stack    []int
	minStack []int
}

func NewMinStack() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64}, //初始化写入一个最大值，用于比较
	}
}

//  入栈操作
func (min *MinStack) Push(x int) {
	//入栈
	min.stack = append(min.stack, x)
	// 取出最小栈的栈顶元素
	top := min.minStack[len(min.minStack)-1]
	// 比较新元素和最小栈的栈顶元素，返回最小值
	minValue := compareValue(x, top)
	// 将最小值插入到最小栈中，
	min.minStack = append(min.minStack, minValue)
}

// 出栈操作
func (min *MinStack) Pop() {
	// 出栈操作，只需要将slice 的len减1即可  保证 minStack 和stack均出栈
	min.stack = min.stack[:len(min.stack)-1]
	min.minStack = min.minStack[:len(min.minStack)-1]
}

// 获取栈顶元素
func (min *MinStack) Top() int {
	return min.stack[len(min.stack)-1]
}

// 获取栈的最小元素
func (min *MinStack) GetMin() int {
	return min.minStack[len(min.minStack)-1]
}

// 比较元素大小
func compareValue(x, y int) int {
	if x < y {
		return x
	}
	return y
}

