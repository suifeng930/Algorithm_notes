package main

func gcd(a, b int) int {

	if a==b {
		return a
	}
	if (a&1)==0 && (b&1)==0 {
		return gcd(a>>1,b>>1)<<1
	}else if (a&1)==0 && (b&1)!=0 {
		return gcd(a>>1,b)
	}else if (a&1)!=0 && (b&1)==0 {
		return gcd(a, b>>1)
	}else {
		big :=Compare2Max(a,b)
		small :=Compare2Min(a,b)
		return gcd(big-small,small)
	}
}
