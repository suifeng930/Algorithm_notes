package main

import "log"

func main()  {

	numbers :=[]int{1,2,3,4,5}
	log.Println(numbers)
	for i :=0;i<10;i++ {
		number := findNearestNumber(numbers)
		printArray(number)
	}
}


