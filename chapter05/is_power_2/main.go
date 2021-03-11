package main

import "log"

func main() {


	log.Println("暴力枚举法")
	log.Println(isPowerOf2V1(32))
	log.Println(isPowerOf2V1(19))

	log.Println("暴力枚举&&移位运算法")
	log.Println(isPowerOf2V2(32))
	log.Println(isPowerOf2V2(19))
	log.Println("按位与运算")
	log.Println(isPowerOf2V3(32))
	log.Println(isPowerOf2V3(19))
}
