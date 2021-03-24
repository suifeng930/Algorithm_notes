package main

import (
	"log"
	"math"
)

func find132pattern(nums []int) bool {
	n := len(nums)
	candidateK := []int{nums[n-1]}
	maxK := math.MinInt64
	log.Println(maxK)
	for i := n - 2; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}
		for len(candidateK) > 0 && nums[i] > candidateK[len(candidateK)-1] {
			maxK = candidateK[len(candidateK)-1]
			candidateK = candidateK[:len(candidateK)-1]
		}
		if nums[i] > maxK {
			candidateK = append(candidateK, nums[i])
		}
	}
	return false
}

func quickSort(nums []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	pivot := partition(nums, startIndex, endIndex)
	quickSort(nums, startIndex, pivot-1)
	quickSort(nums, pivot+1, endIndex)
}

func partition(arr []int, startIndex, endIndex int) int {
	pivot := arr[startIndex]
	left := startIndex
	right := endIndex
	for left != right {
		for left < right && arr[right] > pivot {
			right--
		}
		for left < right && arr[left] <= pivot {
			left++
		}
		if left < right {
			temp := arr[left]
			arr[left] = arr[right]
			arr[right] = temp
		}
	}
	arr[startIndex] = arr[left]
	arr[left] = pivot
	return left
}
