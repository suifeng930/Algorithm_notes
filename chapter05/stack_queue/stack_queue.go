package main

type MyQueue struct {
	inStack  []int
	outStack []int
}

func NewQueue() *MyQueue {
	return &MyQueue{}
}

// 将元素插入到队尾
func (q *MyQueue) enQueue(x int) {
	q.inStack = append(q.inStack, x)
}

// 将元素出队列
func (q *MyQueue) deQueue() int {
	// 如果outStack 为空，则将inStack栈中的数据出栈，并入栈到outStack
	if q.isOutStackEmpty() {
		q.inStackToOutStack()
	}
	val := q.outStack[len(q.outStack)-1]
	// outStack 出栈
	q.outStack = q.outStack[:len(q.outStack)-1]
	return val
}

func (q *MyQueue) inStackToOutStack() {
	for len(q.inStack) > 0 {
		// outStack 入栈
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		// inStack 出栈
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

// 返回当前队头头部元素
func (q *MyQueue) peek() int {
	if q.isOutStackEmpty() {
		q.inStackToOutStack()
	}
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) isEmpty() bool {
	return q.isInStackEmpty() && q.isOutStackEmpty()
}
func (q *MyQueue) isOutStackEmpty() bool {
	return len(q.outStack) == 0
}
func (q *MyQueue) isInStackEmpty() bool {
	return len(q.inStack) == 0
}
