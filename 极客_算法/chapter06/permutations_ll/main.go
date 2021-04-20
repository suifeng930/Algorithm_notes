package main

import "log"

func main() {

	intslice :=[]int{3,3,0,3}

	res := permute(intslice)
	log.Println(res)
}
