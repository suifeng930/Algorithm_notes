package main

//https://leetcode-cn.com/problems/generate-parentheses/
// 括号生成
func generateParenthesis(n int) []string {

	res := []string{}
	// 深度优先遍历
	dfs("", 0, 0, n, &res)
	return res
}

// curStr 当前递归得到的结果
// left 左括号已经用了几个
// right 右括号已经用了几个
// n 左、右括号一共用了几个
// res 生成的结果集
func dfs(curStr string, left, right, n int, res *[]string) {

	// 递归终止条件
	if left == n && right == n {
		*res = append(*res, curStr)
		return
	}
	// 左括号增加
	if left < n {
		dfs(curStr+"(", left+1, right, n, res)
	}
	// 右括号增加
	if right < left {
		dfs(curStr+")", left, right+1, n, res)
	}
}
