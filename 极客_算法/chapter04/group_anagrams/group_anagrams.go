package main



// https://leetcode-cn.com/problems/group-anagrams/
//  字母异位词分组
func groupAnagrams(strs []string) [][]string {


	mp :=map[[26]int][]string{}

	for _,str :=range strs{
		cnt :=[26]int{}
		// 将字符串存到 cnt字符数组中
		for _,b :=range str{
			cnt[b-'a']++
		}
		// 将 字符数组为key str 为value 存入 map中
		mp[cnt]=append(mp[cnt],str)
	}
	// 创建一个二维数组， 遍历map
	ans :=make([][]string,0,len(mp))
	for _,v:=range mp{
		ans =append(ans,v)
	}
	return ans
}