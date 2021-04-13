package main

import "log"

func main() {



	str :=[]string{"eat", "tea", "tan", "ate", "nat", "bat"}

	anagrams := groupAnagrams(str)
	log.Println(anagrams)
}
