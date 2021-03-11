package main

func compareValueToMin(x,y int) int {
	if x<y {
		return x
	}
	return y
}

func compareValueToMax(x,y int) int {
	if x>y {
		return x
	}
	return y
}

func getGreatestCommonDivisor(a, b int) int {
	
	big :=compareValueToMax(a,b)
	small :=compareValueToMin(a,b)
	if big%small==0 {
		return small
	}
	for i :=small/2;i>1;i-- {
		if small%i==0 && big%i==0 {
			return i
		}
	}
	return 1
	
}