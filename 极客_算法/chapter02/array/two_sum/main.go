package main

import "log"

func main() {

	nums :=[]int{2,7,11,15}
	target :=9

	sum := twoSum(nums, target)
	log.Println("return:",sum)

}
