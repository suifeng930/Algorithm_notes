package main



// https://leetcode-cn.com/problems/climbing-stairs/?utm_source=LCUS&utm_medium=ip_redirect&utm_campaign=transfer2china

// 解法1 :递归求解
func climbStairs(n int) int {
	if n==1 {
		return 1
	}
	if n==2 {
		return 2
	}
	return climbStairs(n-1)+ climbStairs(n-2)
}

// 解法2 ：记忆化递归求解   时间复杂度  O(n)  空间复杂度 O(n)
func climbStairsV2(n int) int {

	// 1.初始化一个记忆重复递归子树的数组
	memo :=make([]int,n+1)
	// 2.重复调用递归方法，更新递归子树数组
	return climbStairMemo(n,memo)
}
func climbStairMemo(n int, memo []int) int {
	if memo[n]>0 {
		return memo[n]
	}
	if n==1 {
		memo[n]=1
	}else if n==2 {
		memo[n]=2
	}else {
		memo[n]= climbStairMemo(n-1,memo)+ climbStairMemo(n-2,memo)
	}
	return memo[n]
}

// 动态规划求解     时间复杂度  O(n)  空间复杂度 O(1)
func climbStairV3(n int) int {

	if n==1 {
		return 1
	}
	dp :=make([]int,n+1)
	dp[1]=1
	dp[2]=2
	for i :=3;i<=n;i++ {
		dp[i]=dp[i-1]+dp[n-2]
	}
	return dp[n]
}

// fibonacci  斐波那切数列 求解
func climbStairV4(n int) int {
	if n==1 {
		return 1
	}
	if n==2 {
		return 2
	}
	first :=1
	second :=2
	for i:=3;i<=n;i++ {
		third :=first+second
		first=second  // 更新 f(n-2)的值
		second=third  // 更新 f(n-1)的值
	}
	return second
}