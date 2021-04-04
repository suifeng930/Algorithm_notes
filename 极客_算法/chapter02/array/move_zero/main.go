package main

import "log"

func main() {

	testArr :=[]int{0,1,0,3,12}
	log.Println("before  sort:",testArr)
	//moveZeroesV1(testArr)
	moveZeroV2(testArr)
	log.Println("after  sort:",testArr)

}
