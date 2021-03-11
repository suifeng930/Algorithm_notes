package main

func getGreatestCommonDivisorV3(a, b int) int {

	if a==b {
		return a
	}
	big :=Compare2Max(a,b)
	small:=Compare2Min(a,b)
	return getGreatestCommonDivisorV3(big-small,small)
}

func Compare2Min(a, b int) int {
	if a<b {
		return a
	}
	return b
}
func Compare2Max(a, b int) int {
	if a>b {
		return a
	}
	return b
}
