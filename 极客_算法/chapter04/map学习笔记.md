## 极客-算法学习笔记(08)

[TOC]

### 有效的字母异位词

>给定两个字符串`s`和`t`,编写一个函数来判断`t`是否是`s`的字母异位词
>
>Eg:
>
>输入：s=“anagram”  t=“nagaram”
>
>输出：true
>
>Eg:
>
>输入： s="rat" ,t="car"
>
>输出：false
>
>链接：https://leetcode-cn.com/problems/valid-anagram/



#### 思路(方法1 先排序、后比较)

`t`是`s`的一味词等价于「两个字符串排序后相等」，因此我们可以对字符串`s`和`t`分别排序，看排序后的字符串是否相等即可判断。此外，如果`s`和`t`的长度不同，`t`必然不是`s`的异位词

#### 复杂度分析

* 时间复杂度： O(nlogn). 其中n为字符串长度。排序的时间复杂度为O(nlogn)
* 空间复杂度：O(logn). 排序需要O(logn)的空间复杂度。

#### 代码实现

```go
// 先排序，在比较
func isAnagram(s string, t string) bool {

	s1,s2 :=[]byte(s),[]byte(t)
	sort.Slice(s1, func(i, j int) bool {return s1[i]<s1[j]})
	sort.Slice(s2, func(i, j int) bool {return s2[i]<s2[j]})
	return string(s1)==string(s2)
}
```

#### 思路(采用哈希表)

从另一个角度考虑，`t`是`s`的异位词等价于「两个字符串中字符出现的种类和次数均相等」，由于字符串只包含26个小些字母，因此我们可以维护一个长度为26的频次数组table,先遍历记录字符串`s`中字符串线的频次，然后遍历字符串`t`,减去`table`中对应的频次，如果出现`table[i]<0`,则说明`t`包含一个不在`s`中的额外字符，返回false即可。

#### 复杂度分析

* 时间复杂度：O(n) 其中n为s的长度
* 空间复杂度：O(s) 其中s为字符集大小，此处为s=26

#### 代码实现

```go
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
```

### 字母异位词分组

>给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串
>
>Eg:
>
>输入：`["eat","tea","tan","ate","nat","bat"]`
>
>输出：
>
>```json
>[
>	["ate","eat","tea"],
>	["nat","tan"],
>	["bat"]
>]
>```
>
>**说明：**
>
>- 所有输入均为小写字母。
>- 不考虑答案输出的顺序。
>
>链接：https://leetcode-cn.com/problems/group-anagrams/

#### 思路(计数)

由于护为字母异位词的两个字符串包含的字母相同，因此两个字符串中的相同字母出现的次数一定是相同的，故可以将每个字母出现的次数使用字符串表示，作为哈希表的键。

由于字符串只包含小写字母，因此对于每个字符串，可以使用长度为26的数组记录每个字母出现的次数。

#### 复杂度分析

* 时间复杂度：O(n(k+26))  n是strs中字符串的数量，k是strs 中字符串的最大长度
* 空间复杂度：O(n(k+26))  n是strs中字符串的数量，k是strs 中字符串的最大长度

#### 代码实现

```go
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

```


