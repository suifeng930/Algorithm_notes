package main

import "log"

func main() {

	bitMap := NewBitmapOption()
	bitMap.setBit(126)
	bitMap.setBit(75)
	log.Println("bit map  size:",bitMap.Len())
	log.Println(bitMap.String())
	log.Println(bitMap.getBit(126))
	log.Println(bitMap.getBit(79))
}
