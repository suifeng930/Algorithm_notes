package main

import "log"

func main() {
	str := "aabbcddd"
	arr := strToArr(str)
	log.Println(arr)

}

// 给定一个字符串，将连续字符组合成一个字符串作为数组元素插入到一个数组，最终返回该数组
// eg： 输入str:="aabbcddd"
//      输出["aa","bb","dd"]

func strToArr(str string) []string {

	slow, fast := 0, 0
	result := []string{}
	for fast <= len(str) {
		if fast == len(str) || str[slow] != str[fast] {
			if fast-slow > 1 {
				result = append(result, str[slow:fast])
			}
			slow = fast
		}
		fast++
	}
	return result
}
