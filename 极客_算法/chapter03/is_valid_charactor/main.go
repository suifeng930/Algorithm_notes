package main

import "log"

func main() {


	s := "()[]{}"

	valid := isValid(s)
	log.Println(valid)
}
