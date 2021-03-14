package main

import (
	"fmt"
	"log"
)

func findNearestNumber(numbers []int) []int {

	//1. 从后向前查看逆序区域，找到逆序区域的前一位，也就是数字置换的边界
	index := findTransferPoint(numbers)
	// 如果数字置换边界是0，说明整个数组已经是逆序，无法得到更大的相同数字组成的整数，返回null
	if index==0 {
		return nil
	}
	//2。 把逆序区域的前一位和逆序区域的中刚刚大于它的数字交换位置复制并入参，避免直接修改入参
	changeArray := exchangehead(numbers, index)

	//3.  把原来的逆序区域转为顺序
	val := reverse(changeArray, index)
	return val
}

func findTransferPoint(numbers []int) int {
	for i:= len(numbers)-1;i>0;i-- {
		if numbers[i]>numbers[i-1] {
			return i
		}
	}
	return 0
}

func exchangehead(numbers []int, index int) []int {

	head :=numbers[index-1];
	for i := len(numbers)-1;i>0;i--{
		if head<numbers[i] {
			numbers[index-1]=numbers[i]
			numbers[i]=head
			break
		}
	}
	return numbers
}

func reverse(numbers []int, index int) []int {
	for j:=len(numbers)-1;index<j;j--{
		temp:=numbers[index]
		numbers[index]=numbers[j]
		numbers[j]=temp
		index++
	}
	return numbers
}

func printArray(numbers []int)  {
	for _, number := range numbers {
		fmt.Printf("%d ",number)
	}
	fmt.Println()
	log.Println()
}