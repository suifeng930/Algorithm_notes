package main

//  https://leetcode-cn.com/problems/two-sum/
func twoSum(nums []int,target int) []int {

	// 外层循环控制 x, 内层循环控制num[j]
	for i,x :=range nums{
		for j :=i+1;j<len(nums);j++ {
			// 如果找到符合 a+b=target的两个元素，返回 i，j的数组下标
			if x+nums[j]==target {
				return []int{i,j}
			}
		}
	}
	return nil
}

// 哈希表    空间换时间
// 在遍历的同时，记录一些信息，以省去一层循环，这是 "空间换时间"的想法
// 需要记录已经遍历过的数值和它所对应的下标，可以借助哈希表实现
// 哈希表   key  --> nums的数值，  value --->nums数值的数组下标
func twoSumV2(nums []int, target int) []int {

	//初始化一个map  key 存数组的值，value存数组的下标
	hashTable :=map[int]int{}
	// 遍历数组
	for i,x :=range nums{

		//去hash表中查找 符合 target-x的 key，如果找到就返回，如果没找到，则继续遍历数组
		if j,ok :=hashTable[target-x];ok {
			return []int{j,i}
		}
		// 没有找到，就将当前元素存到hash表中
		hashTable[x]=i
	}
	return nil
}

// 哈希表的第二种写法
func twoSumV3(nums []int, target int) []int {

	result :=[]int{}
	hashTable :=make(map[int]int)
	for i :=0;i< len(nums);i++ {
		if v,exist :=hashTable[target-nums[i]];exist {
			result=append(result,i)
			result=append(result,v)
			return result
		}
		hashTable[nums[i]]=i
	}
	return nil
}