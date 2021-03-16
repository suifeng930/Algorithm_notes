package main

import "strconv"

func compareLength(strA, strB string) int {
	if len(strA) > len(strB) {
		return len(strA)
	}
	return len(strB)
}

func bigNumberSum(bigNumberA, bigNumberB string) string {

	//1. 把两个大整数用数组逆序存储，数组长度等于较大整数位数+1
	maxLength := compareLength(bigNumberA, bigNumberB)

	total := maxLength + 1
	arrayA := make([]int, total)
	arrayB := make([]int, total)
	for i := 0; i < len(bigNumberA); i++ {
		arrayA[i] = int(bigNumberA[len(bigNumberA)-1-i] - '0')
	}
	for i := 0; i < len(bigNumberB); i++ {
		arrayB[i] = int(bigNumberB[len(bigNumberB)-1-i] - '0')
	}
	//2. 构造接收result数组，数组长度等于较大整数位数+1
	result := make([]int, total)
	//3. 遍历数组，按位相加
	for i := 0; i < maxLength+1; i++ {
		temp := result[i]
		temp += arrayA[i]
		temp += arrayB[i]
		//判断是否进位
		if temp >= 10 {
			temp = temp - 10
			result[i+1] = 1
		}
		result[i] = temp
	}
	//4. 把result 数组再次逆序并转成string
	ans := ""
	findFirst := false
	for i := len(result) - 1; i >= 0; i-- {
		if !findFirst {
			if result[i] == 0 {
				continue
			}
			findFirst = true
		}
		ans += strconv.Itoa(result[i])
	}
	if ans == "" {
		return "0"
	}
	return ans
}

func bigNumberSumV2(bigNumberA, bigNumberB string) string {
	// 定义辅助进位标识
	add :=0
	// 接收计算结果
	ans :=""
	for i,j :=len(bigNumberA)-1, len(bigNumberB)-1;i>=0 || j>=0 || add!=0; i,j=i-1,j-1 {
		var x,y int
		if i>=0 {
			x=int(bigNumberA[i]-'0')
		}
		if j>=0 {
			y=int(bigNumberB[j]-'0')
		}
		// 获取每次求和值
		result :=x+y+add
		// 将结果对10求余可巧妙解决进位的问题
		ans=strconv.Itoa(result%10)+ans
		// 将结果相除可初始化进位
		add=result/10
	}
	return ans
}