package main

func findLostNumberV1(nums []int) int {

	table := make(map[int]struct{}, 100)
	//生成一个哈希表
	for i := 1; i <= 100; i++ {
		table[i] = struct{}{}
	}
	//遇到匹配的key 删除对应的key
	for _, val := range nums {
		delete(table, val)
	}
	for i := range table {
		return i
	}
	return 0

}

func findLostNumberV2(nums []int) int {
	//先排序
	quickSort(nums, 0, len(nums)-1)
	//再遍历
	for i := 0; i < len(nums); i++ {
		 temp :=nums[i]
		if temp != (i + 1) {
			return i + 1
		}
	}
	return 0
}

func findLostNumberV3(nums []int) int {

	sum := (100+1)*100>>1
	for _, num := range nums {
		sum=sum-num
	}
	return sum
}

// 异或运算求解
func findLostNumberV4(nums []int) int  {

	sum :=0
	for _, num := range nums {
		sum =sum^num
	}
	return sum
}

func quickSort(arr []int, startIndex, endIndex int) {

	//递归结束条件，startIndex >= endIndex
	if startIndex >= endIndex {
		return
	}
	// 得到基准元素位置
	pivotIndex := partition(arr, startIndex, endIndex)
	// 根据基准元素，分成两部分进行递归排序
	quickSort(arr, startIndex, pivotIndex-1)
	quickSort(arr, pivotIndex+1, endIndex)

}

// 双指针移动
func partition(arr []int, startIndex, endIndex int) int {
	// 获取第1个位置的元素作为基准元素
	pivot := arr[startIndex]
	left := startIndex
	right := endIndex
	for left != right {
		// 控制right指针比较并向左移动
		for left < right && arr[right] > pivot {
			right--
		}
		// 控制left指针比较并向右移动
		for left < right && arr[left] <= pivot {
			left++
		}
		// 交换left 和right指针所指向的元素
		if left < right {
			temp := arr[left]
			arr[left] = arr[right]
			arr[right] = temp
		}
	}
	//pivot 和指针重合点交换
	arr[startIndex] = arr[left]
	arr[left] = pivot
	return left
}
