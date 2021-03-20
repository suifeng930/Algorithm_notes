package main

type Bitmap struct {
	size int
	words []int64
}

func NewBitmapOption( size int) *Bitmap  {
	words :=make([]int64,getWordIndex(size-1)+1)
	return &Bitmap{
		size:  size,
		words: words,
	}
}

// 定位bitmap某一位所对应的word
func getWordIndex(bitIndex int) int  {
	// 右移6位，相当于除以64
	return bitIndex>>6
}

// 判断bitMap某一位的状态
func (b *Bitmap)getBit(bitIndex int) bool {
	if bitIndex<0 || bitIndex>b.size-1 {
		panic ("超出了bitmap有效范围")
	}
	wordIndex :=getWordIndex(bitIndex)
	return (b.words[wordIndex] & (1 <<bitIndex)) !=0
}

// 把bitMap某一位设置为true
func (b *Bitmap)setBit(bitIndex int)  {
	if bitIndex<0 || bitIndex>b.size-1 {
		panic ("超出了bitmap有效范围")
	}
	wordIndex :=getWordIndex(bitIndex)
	b.words[wordIndex] |= 1 <<bitIndex
}

