package main

import "log"

func main() {


	arr :=[]int{0,0,1,1,1,2,2,3,3,4}
	printArray(arr)
	duplicatesLen := removeDuplicates(arr)
	log.Println(duplicatesLen)
	printArray(arr)
}
