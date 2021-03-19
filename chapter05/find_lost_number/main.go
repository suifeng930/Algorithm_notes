package main

import "log"

func main() {

	array := createArray(100)
	lostNumber := findLostNumberV1(array)
	log.Println(lostNumber)
	v2 := findLostNumberV2(array)
	log.Println(v2)
	v3 := findLostNumberV3(array)
	log.Println(v3)

	arr :=[]int{3,1,3,2,4,1,4}
	v4 := findLostNumberV4(arr)
	log.Println(v4)

}

func createArray(length int) []int {
	array :=make([]int,length)
	for i :=0;i<length;i++ {
		if i==78 {
			continue
		}
		array[i]=i+1
	}
	for i :=0;i< len(array);i++ {
		if array[i]==0 {
			array=append(array[:i],array[i+1:]...)
		}
	}
	return array
}
