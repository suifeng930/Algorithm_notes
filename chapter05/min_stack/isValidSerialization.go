package main

import "log"

func isValidSerialization(preorder string) bool {

	n := len(preorder)
	stack := []int{1}
	for i := 0; i < n; {
		if len(stack) == 0 {
			return false
		}
		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' {
			stack[len(stack)-1]--
			if stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			}
			i++
		} else {
			//读一个数字
			for i < n && preorder[i] != ',' {
				i++
			}
			stack[len(stack)-1]--
			if stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, 2)
		}
	}
	return len(stack) == 0
}

func isValidSerializationV2(preorder string) bool {
	n := len(preorder)

	slots := 1
	for i := 0; i < n; {
		if slots == 0 {
			return false
		}
		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' {
			slots--
			i++
		} else {
			// 读到的是数字
			for i < n && preorder[i] != ',' {
				i++
			}
			slots++ // slots =slots-1+2
		}
	}
	return slots == 0
}

func isValidSerializationV3(preorder string) bool {

	degree := 1
	n := len(preorder)
	for i := 0; i < n; {
		if degree == 0 {
			return false
		}
		val :=preorder[i]
		log.Println(val)
		if preorder[i] == ',' {
			i++
		} else if preorder[i] == '#' { //空叶子节点
			degree -= 1
			i++
		} else {
			for i <n && preorder[i]!=',' {
				i++
			}
			degree += 1 //非叶子节点
		}
	}
	return degree == 0
}
