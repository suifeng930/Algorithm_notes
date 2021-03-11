package main

func CompareToMin(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func CompareToMax(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func getGreatestCommonDivisorV2(a, b int) int {
	big :=CompareToMax(a,b)
	small :=CompareToMin(a,b)
	if big%small==0 {
		return small
	}
	return getGreatestCommonDivisorV2(big%small,small)

}