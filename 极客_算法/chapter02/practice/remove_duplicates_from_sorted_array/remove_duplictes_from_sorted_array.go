package main

import "fmt"

//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums)==0 {
		return 0
	}
	slow :=0
	for fast:=1;fast< len(nums);fast++ {
		if nums[fast]!=nums[slow] {
			slow++
			nums[slow]=nums[fast]
		}
	}
	return slow+1
}

func printArray(nums []int)  {

	for i :=0;i< len(nums);i++{
		fmt.Printf("%d-->",nums[i])
	}
	fmt.Println()

}
