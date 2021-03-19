package main

func quickSort(arr []int,startIndex ,endIndex int)  {

	//递归结束条件，startIndex >= endIndex
	if startIndex>=endIndex{
		return
	}
	// 得到基准元素位置
	pivotIndex :=Partition(arr, startIndex,endIndex)
	// 根据基准元素，分成两部分进行递归排序
	quickSort(arr,startIndex,pivotIndex-1)
	quickSort(arr,pivotIndex+1,endIndex)

}

// 双指针移动
func partition(arr []int, startIndex, endIndex int) int {
	// 获取第1个位置的元素作为基准元素
	pivot :=arr[startIndex]
	left :=startIndex
	right :=endIndex
	for left!=right {
		// 控制right指针比较并向左移动
		for left<right && arr[right]>pivot {
			right--
		}
		// 控制left指针比较并向右移动
		for left<right && arr[left]<=pivot {
			left++
		}
		// 交换left 和right指针所指向的元素
		if left<right {
			temp :=arr[left]
			arr[left]=arr[right]
			arr[right]=temp
		}
	}
	//pivot 和指针重合点交换
	arr[startIndex]=arr[left]
	arr[left]=pivot
	return left
}

func Partition(arr []int, startIndex, endIndex int) int {

	// 取第一个元素作为基准元素
	pivot :=arr[startIndex]
	mark :=startIndex
	// 从基准元素的下一个位置开始遍历数组
	for i :=startIndex+1;i<=endIndex;i++{
		// 如果遍历到的元素小于基准元素，则mark++，
		//并将最新遍历到的元素和mark指针所在位置的元素交换位置
		if arr[i]<pivot {
			mark++
			temp :=arr[mark]
			arr[mark]=arr[i]
			arr[i]=temp
		}
	}
	arr[startIndex]=arr[mark]
	arr[mark]=pivot
	return mark

}