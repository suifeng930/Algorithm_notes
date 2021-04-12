package main

func isValid(s string) bool {

	// 1. 先判断字符串是否合法
	n := len(s)
	if n%2 == 1 {
		return false
	}
	// 2.初始化一个map 哈希表的键为右括号，值为相同类型的左括号。
	compare := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if compare[s[i]] > 0 {
			// 3. 如果栈为空，或者栈顶元素和map中的value值不匹配，返回false
			if len(stack) == 0 || stack[len(stack)-1] != compare[s[i]] {
				return false
			}
			//出栈
			stack = stack[:len(stack)-1]
		} else { // 4.如果value 是左括号，则入栈
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func isValidV2(s string) bool {
	n := len(s)

	// 如果字符串长度不为偶数，直接返回
	if n%2 != 0 {
		return false
	}
	// 定一个map
	compare := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}

	for i := 0; i < n; i++ {
		// 如果当前字符传是右括号，则将栈顶取出，和map中的元素匹配，如果匹配成功，栈顶元素出栈
		if compare[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != compare[s[i]] {
				return false
			}
			//栈顶元素出栈，找到匹配的括号对
			stack = stack[:len(stack)-1]
		} else {
			//如果是左括号，则将括号入栈
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
