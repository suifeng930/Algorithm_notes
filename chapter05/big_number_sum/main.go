package main

import "log"

func main() {


	strA :="426709752318"
	strB :="95481253129"
	sum := bigNumberSum(strA, strB)
	log.Println(sum)

	v2 := bigNumberSumV2(strA, strB)
	log.Println(v2)

}
