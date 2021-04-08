package main

import (
	"sort"
)

//  https://leetcode-cn.com/problems/merge-sorted-array/
// 合并两个有序数组
func merge(nums1 []int,m int,nums2 []int ,n int)  {

	// 将nums2 合并到nums1的尾部
	copy(nums1[m:],nums2)
	// 对num1排序
	sort.Ints(nums1)
}


// 双指针解法，   创建了额外的内存空间
func mergeV2(nums1 []int,m int,nums2 []int ,n int)  {

	//初始化排序数组
	sorted :=make([]int,0,m+n)
	p1,p2 :=0,0
	for  {
		// 如果 nums1 为空，返回nums2
		if p1==m {
			sorted=append(sorted,nums2[p2:]...)
			break
		}
		// 如果nums2为空，返回nums1
		if p2==n {
			sorted=append(sorted,nums1[p1:]...)
			break
		}
		// 如果nums1[i]<nums2[i] 将nums1[i] 放到排序后的数组中
		if nums1[p1]<nums2[p2] {
			sorted=append(sorted,nums1[p1])
			p1++
		}else {
			sorted=append(sorted,nums2[p2])
			p2++
		}
	}
	// 将排序后的数组，复制到原nums1数组中
	copy(nums1,sorted)
}

// 双指针解法，从后向前求解
func mergeV3(nums1 []int,m int,nums2 []int ,n int)  {

	// 因为两个数组是有序数组，比较两个数组的尾部元素大小，将尾数大的元素追加到nums1的尾部
	for p1 ,p2 ,tail :=m-1,n-1,m+n-1;p1>=0 ||p2>=0; tail--{
		var curr int
		if p1==-1 {
			curr=nums2[p2]
			p2--
		}else if p2==-1 {
			curr=nums1[p1]
			p1--
			//如果nums1尾部大于nums2尾部，将nums1元素追加到 最后  tail--
		}else if nums1[p1]>nums2[p2] {
			curr=nums1[p1]
			p1--
		}else {
			curr=nums2[p2]
			p2--
		}
		nums1[tail]=curr
	}
}