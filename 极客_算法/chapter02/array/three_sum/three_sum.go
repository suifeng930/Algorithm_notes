package main

// https://leetcode-cn.com/problems/3sum/
// 快速排序+双指针解法
func threeSum(nums []int) [][]int {

	// length
	ans := make([][]int, 0)
	length := len(nums)
	if nums == nil || length < 3 {
		return ans
	}
	//1. 使用快速排序对数组排序
	quickSort(nums, 0, length-1)

	//sort.Ints(nums)

	//2. 使用双指针，左右收敛求解
	for i := 0; i < length-2; i++ { // length -2 少两次遍历
		if nums[i] > 0 {
			break // 如果当前k  大于零 则三数之和一定大于0，因此结束当前循环
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue // 已经统计过了 nums[i] ,去重
		}
		leftIndex := i + 1
		rightIndex := length - 1
		for leftIndex < rightIndex {
			sum := nums[i] + nums[leftIndex] + nums[rightIndex]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[leftIndex], nums[rightIndex]})
				for leftIndex < rightIndex && nums[leftIndex] == nums[leftIndex+1] {
					leftIndex++ // 存在相同元素，跳过所有重复对左元素，去重
				}
				for leftIndex < rightIndex && nums[rightIndex] == nums[rightIndex-1] {
					rightIndex-- //跳过所有重复的右元素去重
				}
				leftIndex++
				rightIndex--
			} else if sum < 0 { //说明左边数值太小了，需要右移动
				leftIndex++
			} else if sum > 0 { //说明右边数值太大了，需要左移动
				rightIndex--
			}
		}
	}
	return ans
}

func quickSort(nums []int, startIndex, endIndex int) {

	if startIndex >= endIndex {
		return
	}
	pivot := partition(nums, startIndex, endIndex)
	quickSort(nums, startIndex, pivot-1)
	quickSort(nums, pivot+1, endIndex)

}

func partition(nums []int, startIndex, endIndex int) int {

	pivotValue := nums[startIndex]
	left := startIndex
	right := endIndex
	for left != right {
		for left < right && nums[right] > pivotValue {
			right--
		}
		for left < right && nums[left] <= pivotValue {
			left++
		}
		if left < right {
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
		}
	}
	nums[startIndex] = nums[left]
	nums[left] = pivotValue
	return left

}
