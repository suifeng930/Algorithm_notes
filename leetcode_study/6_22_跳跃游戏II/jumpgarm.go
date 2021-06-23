package main

import "log"

func main() {

	arr := []int{2, 1, 0, 1, 4}
	end := canJump(arr)
	log.Println(end)
}

//给定一个非负整数数组，你最初位于数组的第一个位置。
//
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
//你的目标是使用最少的跳跃次数到达数组的最后一个位置。
//
//假设你总是可以到达数组的最后一个位置。
//
//
//
//示例 1:
//
//输入: [2,3,1,1,4]
//输出: 2
//解释: 跳到最后一个位置的最小跳跃数是 2。
//     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
//示例 2:
//
//输入: [2,3,0,1,4]
//输出: 2

func canJump(nums []int) int {

	end := 0
	jumpFlag := 0
	steps := 0
	for i, val := range nums[:len(nums)-1] { // 最后一个元素可不用访问
		jumpFlag = compareToMax(jumpFlag, i+val) // 获取最大跳跃步数
		if i == end {                            // flag可达，steps++
			end = jumpFlag
			steps++
		}
	}
	return steps
}
func compareToMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
