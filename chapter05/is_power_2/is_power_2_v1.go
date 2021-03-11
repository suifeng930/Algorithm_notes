package main

func isPowerOf2V1(num int) bool {

	temp :=1
	for temp<=num {
		if temp==num {
			return true
		}
		temp=temp*2
	}
	return false
}
