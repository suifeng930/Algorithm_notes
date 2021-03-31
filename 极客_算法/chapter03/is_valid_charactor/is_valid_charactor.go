package main

func isValid(s string) bool {

	// 1. 先判断字符串是否合法
	n :=len(s)
	if n%2==1 {
		return false
	}
	// 2.初始化一个map 哈希表的键为右括号，值为相同类型的左括号。
	compare :=map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}
	stack :=[]byte{}
	for i :=0;i<n;i++ {
		if compare[s[i]]>0 {
			// 3. 如果栈为空，或者栈顶元素和map中的value值不匹配，返回false
			if len(stack)==0 || stack[len(stack)-1]!=compare[s[i]] {
				return false
			}
			//出栈
			stack=stack[:len(stack)-1]
		}else {// 4.如果value 是左括号，则入栈
			stack=append(stack,s[i])
		}
	}
	return len(stack)==0
}