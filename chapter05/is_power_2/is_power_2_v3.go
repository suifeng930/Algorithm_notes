package main

//按位与运算 巧解 n&(n-1)==0
func isPowerOf2V3(num int) bool {
	return (num&(num-1))==0

}
