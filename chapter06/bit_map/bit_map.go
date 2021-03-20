package main

import (
	"bytes"
	"fmt"
)

type Bitmap struct {
	size int
	words []uint64
}

func NewBitmapOption( ) *Bitmap  {
	return &Bitmap{}
}

// 定位bitmap某一位所对应的word
func getWordIndex(bitIndex int) int  {
	// 右移6位，相当于除以64
	return bitIndex>>6
}

// 判断bitMap某一位的状态
func (b *Bitmap)getBit(bitIndex int) bool {

	wordIndex :=getWordIndex(bitIndex)
	bit :=uint(bitIndex%64)
	return wordIndex< len(b.words) && (b.words[wordIndex] & (1 <<bit)) !=0
}

// 把bitMap某一位设置为true
func (b *Bitmap)setBit(bitIndex int)  {
	word :=getWordIndex(bitIndex)
	bit :=uint(bitIndex%64)
	for word>=len(b.words) {
		b.words=append(b.words,0)
	}
	//判断bitIndex 是否已经存在bitmap中
	if b.words[word] & 1 <<bit==0{
		b.words[word] |= 1<<bit
		b.size++
	}
}

func (b *Bitmap) Len() int {
	return b.size
}


func (b *Bitmap) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, v := range b.words {
		if v == 0 {
			continue
		}
		for j := uint(0); j < 64; j++ {
			if v&(1<<j) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*uint(i)+j)
			}
		}
	}
	buf.WriteByte('}')
	fmt.Fprintf(&buf,"\nLength: %d", b.size)
	return buf.String()
}