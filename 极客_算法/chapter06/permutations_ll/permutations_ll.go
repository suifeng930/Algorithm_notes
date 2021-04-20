package main

import (
	"log"
	"sort"
)

// https://leetcode-cn.com/problems/permutations-ii/
// 46 全排列
// 递归决策树
func permute(nums []int) [][]int {

	n := len(nums)
	res :=[][]int{}
	if n==0 {
		return res
	}
	// 将数组排序
	sort.Ints(nums)
	//  存储全排列的一组数据 path 和used是全程共享的，一定要注意恢复状态
	path :=[]int{}
	// 标记被扫描到的元素，如果被扫描到记为true
	used :=make([]bool,n)
	dfs(nums,0,n,path,used,&res)
	return res
}

func dfs(nums []int, depth int, length int, path []int, used []bool, res *[][]int) {

	// 如果 深度和数组相同 退出递归条件
	if depth==length {
		temp :=make([]int,len(path)) // 保存path的快照
		copy(temp,path)
		*res=append(*res,temp)
		log.Println(path)
		return
	}
	for i :=0;i<length;i++{
		if used[i]  {
			continue
		}
		// 比较相邻两个元素，如果两个数重复，去重，因为经过排序后的数组从左到右一定是有序
		if i>0 && used[i-1] && nums[i]==nums[i-1] {
			continue
		}
		path=append(path,nums[i])
		used[i]=true // 更改当前状态，递归调用下一层
		dfs(nums,depth+1,length,path,used,res)
		path=path[:len(path)-1] // 每次返回后恢复状态，供下一层使用
		used[i]=false
	}
}
