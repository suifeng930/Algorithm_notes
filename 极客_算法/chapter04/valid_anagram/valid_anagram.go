package main

import "sort"

// https://leetcode-cn.com/problems/valid-anagram/
// 先排序，在比较
func isAnagram(s string, t string) bool {

	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

// 创建hash table
func isAnagramV2(s string, t string) bool {

	var c1, c2 [26]int
	//
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}

// 创建map 一个字节map
func isAnagramV3(s, t string) bool {

	// 如果两个字符串长度不相等直接返回false
	if len(s)!=len(t) {
		return false
	}

	// 初始化一个unicode编码的map
	cnt:=map[rune]int{}

	// 遍历字符串s,将每一个字符都存到map中
	for _,ch :=range s{
		cnt[ch]++
	}
	// 遍历字符串t,将每一个字符都去map中查找，如果查到key,则删除该元素
	for _,ch :=range t{
		cnt[ch]--
		if cnt[ch]<0 {
			return false
		}
	}
	return true
}