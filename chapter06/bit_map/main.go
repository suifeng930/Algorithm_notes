package main

import "log"

func main() {

	bitMap := NewBitmapOption(128)
	bitMap.setBit(126)
	bitMap.setBit(75)
	log.Println(bitMap.getBit(126))
	log.Println(bitMap.getBit(75))
}
