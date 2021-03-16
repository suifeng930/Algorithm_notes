package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println(removeKDigits("1593212",3))
	log.Println(removeKDigits("30200",1))
	log.Println(removeKDigits("10",2))
	log.Println(removeKDigits("541270936",3))

	S :="sterea"
	for _, i := range S {
		fmt.Printf("%T %+v \n",i,i)

	}
	for i :=0;i<len(S);i++ {
		fmt.Printf("%T %+v \n",i,S[i])

	}
}
