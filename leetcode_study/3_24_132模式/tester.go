package main

import "log"

func main() {

	//val :=[]int{3,1,4,2}
	//val :=[]int{-1,3,2,0}
	val :=[]int{1,0,1,-4,-3}
	pattern := find132pattern(val)
	log.Println(pattern)

}
