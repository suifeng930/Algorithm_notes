package main

func isPowerOf2V2(num int) bool {
	temp :=1
	for temp<=num {
		if temp==num {
			return true
		}
		temp=temp<<1
	}
	return false
}