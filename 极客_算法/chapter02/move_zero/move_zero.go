package main

func moveZeroesV1(nums []int) {

	lastNonZeroIndex := 0
	// 先将非零元素移动到前面
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZeroIndex] = nums[i]
			lastNonZeroIndex++
		}
	}
	// 再将后面的位置进行0填充
	for i := lastNonZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroV2(nums []int)  {

	// 双指针解法

	leftIndex ,rightIndex,n :=0,0, len(nums)
	// 移动右指针
	for rightIndex<n {
		//如果右指针 所在的元素不等于0
		if nums[rightIndex]!=0 {
			// 交换左右指针的元素
			nums[leftIndex],nums[rightIndex] =nums[rightIndex],nums[leftIndex]
			// 左指针向右移动
			leftIndex++
		}
		// 右指针右移
		rightIndex++
	}

}