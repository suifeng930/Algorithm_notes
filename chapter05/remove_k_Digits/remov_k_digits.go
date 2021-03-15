package main

import "strings"

func removeKDigits(num string, k int) string {

	//初始化一个栈，用于接收所有的数字
	stack :=[]byte{}
	//遍历原数列
	for i := range num {
		//遍历当前数字
		digit :=num[i]
		// 如果k>0 且 栈顶数字大于遍历到的当前数字，
		for k>0 && len(stack)>0 && digit<stack[len(stack)-1] {
			// 栈顶数字出栈
			stack=stack[:len(stack)-1]
			// k数字减1
			k--
		}
		// 遍历到的当前数字入栈
		stack=append(stack,digit)
	}
	// 获取截取之后的栈
	stack=stack[:len(stack)-k]

	// 删除数字中包含前导为0的数字
	digits :=strings.TrimLeft(string(stack),"0")
	// 如果最终的字符串为 空 返回 0
	if digits=="" {
		digits="0"
	}
	return digits
}