package main

import "log"

func main() {

	arr :=[]int{4,4,6,5,3,2,8,1}
	quickSort(arr,0, len(arr)-1)
	log.Println(arr)
	
}
