package main

import "log"

func main() {

	log.Println("暴力枚举法求解：")
	divisor := getGreatestCommonDivisor(25, 5)
	log.Println(divisor)
	log.Println(getGreatestCommonDivisor(100,80))
	log.Println(getGreatestCommonDivisor(27,14))

	log.Println("辗转相除法求解：")
	divisorV2 := getGreatestCommonDivisorV2(25, 5)
	log.Println(divisorV2)
	log.Println(getGreatestCommonDivisorV2(100,80))
	log.Println(getGreatestCommonDivisorV2(27,14))
	log.Println("更相减损求解：")
	divisorV3 := getGreatestCommonDivisorV3(25, 5)
	log.Println(divisorV3)
	log.Println(getGreatestCommonDivisorV3(100,80))
	log.Println(getGreatestCommonDivisorV3(27,14))
	log.Println(getGreatestCommonDivisorV3(23456,21))

	log.Println("更相减损&&移位运算求解：")
	divisorV4 := gcd(25, 5)
	log.Println(divisorV4)
	log.Println(gcd(100,80))
	log.Println(gcd(27,14))
	log.Println(gcd(23456,21))
}
