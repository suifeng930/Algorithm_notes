package main

// https://leetcode-cn.com/problems/container-with-most-water/
// 盛最多水的容器
func maxArea(height []int) int {

	maxArea := 0
	leftIndex, rightIndex := 0, len(height)-1
	// 双指针，左指针不大于右指针
	for leftIndex < rightIndex {
		// 获取到左右指针的值
		heiLeft, heiRight := height[leftIndex], height[rightIndex]
		//计算 area
		area := compare2Min(heiLeft, heiRight) * (rightIndex - leftIndex)
		// 比较并赋值给最大area
		if maxArea < area {
			maxArea = area
		}
		if heiLeft < heiRight {
			// 如果左边 height 小于右边height 左指针右移动
			leftIndex++
		} else {
			// 反之右指针左移动
			rightIndex--
		}
	}
	return maxArea
}

func compare2Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func compare2Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
