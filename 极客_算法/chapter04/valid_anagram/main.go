package main

import "log"

func main() {


	s := "anagram"
	t := "nagaram"

	//anagram := isAnagram(s, t)
	//anagram := isAnagramV2(s, t)
	anagram := isAnagramV3(s, t)
	log.Println(anagram)
}
