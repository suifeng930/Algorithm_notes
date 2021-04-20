package main

// https://leetcode-cn.com/problems/combinations/
// 77. 组合  递归求解
func combine(n int, k int) [][]int {

	res :=[][]int{}
	temp :=[]int{}
	for i :=1;i<=k;i++ {
		temp=append(temp,i)
	}
	temp=append(temp,n+1)
	for j:=0;j<k; {
		comb :=make([]int,k)
		copy(comb,temp[:k])
		res =append(res,comb)
		for j=0;j<k&& temp[j]+1==temp[j+1];j++ {
			temp[j]=j+1
		}
		temp[j]++
	}
	return res

}



// 77. 组合  递归求解
func combineV2(n int, k int) [][]int {

	res := [][]int{}
	var helper func(n, k int, path []int)
	helper = func(n, k int, path []int) {
		if n < k || k == 0 { //
			if k == 0 {
				temp := make([]int, len(path))
				copy(temp, path)
				res = append(res, temp)
			}
			return
		}
		helper(n-1, k-1, append(path, n))
		helper(n-1, k, path)
	}
	helper(n, k, []int{})
	return res

}



// 77. 组合  递归求解
func combineV3(n int, k int) [][]int {

	res :=[][]int{}
	path:=[]int{}
	dfs(n,k,path,&res)
	return res

}

func dfs(n,k int,path []int,res *[][]int)  {

	if n<k {
		return
	}
	if k==0 {
		r :=make([]int,len(path))
		copy(r,path)
		*res=append(*res,r)
		return
	}

	for n>0 {
		path=append(path,n)
		n--
		dfs(n,k-1,path,res)
		path=path[:len(path)-1]
	}
}

