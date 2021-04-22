package main

import "log"

//  https://leetcode-cn.com/problems/subsets/

// 子集  递归求解
func subsets(nums []int) [][]int {

	res :=[][]int{}
	path :=[]int{}
	length := len(nums)
	dfs(0,length,path,nums,&res)
	return res
}

func dfs(depth, length int, path, nums []int, res *[][]int) {

	// 递归终止条件
	if depth==length {
		temp:=make([]int, len(path))
		copy(temp,path)
		log.Println(path)
		*res=append(*res,path)
		return
	}
	// 考虑选择当前位置元素
	path=append(path,nums[depth])
	dfs(depth+1,length,path,nums,res)
	path=path[:len(path)-1]
	// 不考虑当前位置元素
	dfs(depth+1,length,path,nums,res)
}