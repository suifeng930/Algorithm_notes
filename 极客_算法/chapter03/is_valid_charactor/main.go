package main

import "log"

func main() {


	s := "()[]{}"

	//valid := isValid(s)
	valid := isValidV2(s)
	log.Println(valid)
}
