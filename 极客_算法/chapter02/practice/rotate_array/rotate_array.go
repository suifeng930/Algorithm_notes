package main


//  https://leetcode-cn.com/problems/rotate-array/
// 旋转数组
func rotate(nums []int, k int)  {
	// 创建一个空数组
	temp :=make([]int,len(nums))
	//遍历原数组， 将原数组下标为 ii 的元素放至新数组下标为 (i+k)mod n的位置
	for i :=0;i<len(nums);i++{
		temp[(i+k)%len(nums)]=nums[i]
	}
	// 将temp数组copy到原数组
	copy(nums,temp)
}

// 翻转数组
func rotateV2(nums []int,k int)  {
	k %= len(nums)
	// 1.先对整个数组实行翻转，这样原数组中需要翻转的子数组，就会跑到数组最前端
	reverse(nums)
	// 2.以k为分界线，将做翻转后的数组分割为左右两个数组，先翻转[0,k]的数组，
	reverse(nums[:k])
	// 3.翻转[k:n]的数组
	reverse(nums[k:])
}

func reverse(a []int)  {

	for i,n :=0,len(a);i<n/2;i++ {
		a[i],a[n-i-1]=a[n-1-i],a[i]
	}
}