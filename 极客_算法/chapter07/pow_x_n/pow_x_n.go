package main


// https://leetcode-cn.com/problems/powx-n/
//  50 实现x的n次幂函数
//  递归分治
func myPow(x float64, n int) float64 {

	// 如果n 为负数，将其转换为 1/  x^n
	if n<0 {
		x =1/x
		n=-n
	}
	return dfs(x,n)
}

func dfs(x float64, n int) float64 {

	// 递归终止条件
	if n==0 {
		return 1.0
	}
	// 当前层逻辑
	res := dfs(x, n/2)
	if n%2==0 {
		// n为偶数
		return res*res
	}else {
		// n为奇数
		return res*res*x
	}
}