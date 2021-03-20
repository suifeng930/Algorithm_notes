## 漫画算法-小灰的算法之旅（29）

[TOC]

### Bitmap的秘密

#### BitMap 如何做到多维交叉运算的

Bit即比特，是目前计算机系统里面数据的最小单位，8个bit即为一个Byte。一个bit的值，或者是0，或者是1；也就是说一个bit能存储的最多信息是2.

Bitmap可以理解为通过一个bit数组来存储特定数据的一种数据结构；由于bit是数据的最小单位，所以这种数据结构往往是非常节省存储空间的。例如一个公式有8个员工，现在需要记录公司人员的考勤记录，传统的方案是记录下每天正常考勤的员工的ID列表，比如 2012-01-01 :[1,2,3,4,5,6,7,8]。假如员工ID采用Byte数据类型，则每天的考勤记录需要存8个byte。如果公司人员为n呢？那就需要n个byte的存储空间了。如果采用另一种方案： 构造一个8bit（01110011）的数组，将这8个员工跟员工工号分别映射到这个8个位置，如果当前正常考勤了，则将对应的这个位置置为1，否则置为0；这样每天采用恒定的一个byte即可保存当天的考勤记录。

综上所述，Bitmap节省大量的存储空间，因此可以被一次加载到内存中。再看其结构的另一个更重要的特点，它也显现出巨大威力：**就是很方便通过位的运算(AND/OR/XOR/NOT)，高效的对多个Bitmap数据进行处理**，这点很重要，它直接的支持了多维交叉计算能力。比如上面员工考勤的例子里，如果想要知道哪个员工最近两天都没有来，只要将昨天的bitmap和今天的bitmap做一个按位与"OR"运算，然后检查哪些位置是0的，就可以得到最近两天都没有来的员工数据信息了；例如：

```sh

# 01110011    昨天的bitmap
# 01101010    今天的bitmap
# 按位与运算 --> 01111011   --> bitmap中为0 的数位代表的员工即为缺勤的信息。
```

#### BitMap如何做到高速运算的

回忆一下前面，浪费的有两个部分：

* 其一是存储空间的浪费
* 其二是计算资源的浪费

**存储空间的浪费**：BitMap比文件强多了，但是仍然有浪费的嫌疑。它需要保存到外部存储(数据库或文件)，计算时需要从外部存储加载到内存，因此存储的BitMap越大，需要的外部存储空间就越大；并且计算时`I/O`的消耗会更大，加载BitMap的时间也就越长。

计算资源的浪费: 计算时要加载到内存，越大的Bitmap消耗的内存越多；位数越多，计算时消耗的cpu时间也就越多。

对于第一种浪费，最直觉的方案就是可以引入一些文件压缩技术，比如`gzip/lzo`之类的,对存储的BitMap文件进行压缩，再加载BitMap的时候再进行解压，这样可以很好的解决存储空间的浪费，以及加载时`I/O`的消耗；代价则是压缩/解压缩都需要消耗更多的CPU/内存资源；并且文件压缩技术对第二种浪费也无能为力。因此只有系统有足够多空闲的CPU资源而`I/O`成为瓶颈的情况下，可以考虑引入文件压缩技术。

那么有没有一些技术可以同时解决这两种浪费呢？有的，常见的压缩技术都是基于RLE(Run Length Encoding)详情：http://en.wikipedia.org/wiki/Run-length_encoding

RLE 编码很简单，比较适合有很多连续字符的数据，比如以下边的BitMap为例：

```sh
# 00000000110000000000010011100000000000
# 可以编码为0,8,2,11,1,2,3,11
```

其意思是：第1位为0，连续有8个零，接下来是2个1，11个0，一个1，2个0，3个1，最后是11个0(当然此处知识对RLE的基本原理解释，实际应用中的编码并不完全是这样的)。

可以预见，对于一个很大的BitMap，如果里面的数据分布很稀疏(说明有很多大片连续的0)，采用RLE编码后，占用的空间会比原始的BitMap小很多。

同时引入一些对齐的技术，可以让采用RLE编码的BitMap不需要进行解压缩，就可以直接进行`AND/OR/XOR`等各类计算;因此采用这类压缩技术的Bitmap，加载到内存后还是以压缩的方式存在，从而可以保证计算时候的低内存消耗；而**采用word（计算机的字长，64位系统就是64bit）对齐等计算又保证了CPU资源的高效利用**。因此采用这类压缩技术的bitMap，保持了Bitmap数据结构最重要的一个特性，就是高效的正对每一个bit的逻辑运算。

常见的压缩技术包括 BBC（有专利保护）：

WAH（[ http://code.google.com/p/compressedbitset/）](http://code.google.com/p/compressedbitset/）)

和 EWAH([ http://code.google.com/p/javaewah/ ](http://code.google.com/p/javaewah/))。在 Apache Hive 里边使用了 EWAH。

#### BitMap代码实现

```java

// 每一个word是一个long类型元素，对应一个64位二进制数据
private long[] words;
//BitMap 的位数大小
private int size;

public MyBitmap(int size){
  this.size=size;
  this.words=new long[(getWordIndex(size-1)+1)];
}

/**
*  判断bitmap某一位的状态
*  param bitIndex. 位图的第 bitIndex 位
*/
public boolean getBit(int bitIndex){
  if(bitIndex<0 || bitIndex>size-1){
    throw new IndexOutOfBoundsException("超过Bitmap有效范围");
  }
  int wordIndex=getWordIndex(bitIndex);
  return (words[wordIndex] & (1L<<bitIndex))!=0;
}

/**
* 把bitmap某一位设置为true
* param bitIndex 位图的第bitIndex 位
*/
public void setBit(int bitIndex){
  if(bitIndex<0 || bitIndex>size-1){
    throw new IndexOutOfBoundsException("超过Bitmap有效范围");
  }
  int wordIndex=getWordIndex(bitIndex);
  words[wordIndex] |=(1L<<bitIndex);
}

/**
* 定位bitMap某一位所对应的word
* param bitIndex 位图的第 bitIndex 位
*/
private int getWordIndex(int bitIndex){
  // 右移6位，相当于除以64
  return bitIndex >>6;
}

public static void main(String[] args){
  MyBitmap bitMap=new MyBitmap(128);
  bitMap.setBit(126);
  bitMap.setBit(75);
  System.out.println(bitMap.getBit(126));
  System.out.println(bitMap.getBit(78));
}
```

在上述代码中，使用一个命名为words的long类型数组来存储所有的二进制位。每一个long元素占用其中的64位。

如果要把Bitmap的某一位设为1，需要进过两步：

* 定位到words中的对应的long元素
* 通过与运算修改long的值

如果要查看Bitmap的某一位是否为1，也需要经过两步：

* 定位到words中对应的long元素
* 判断long元素的对应的二进制位是否为1

#### golang实现

```go
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

func main() {

	bitMap := NewBitmapOption(128)
	bitMap.setBit(126)
	bitMap.setBit(75)
	log.Println(bitMap.getBit(126))
	log.Println(bitMap.getBit(75))
}
```



### 关于BitMap算法一些处理大数据问题的场景：

（1）给定40亿个不重复的 int的整数，没排过序的，然后再给一个数，如何快速判断这个数是否在那40亿个数当中。

解法：遍历40亿数字，映射到BitMap中，然后对于给出的数，直接判断指定的位上存在不存在即可。

（2）使用位图法判断整形数组是否存在重复

解法：遍历一遍，存在之后设置成1，每次放之前先判断是否存在，如果存在，就代表该元素重复。

（3）使用位图法进行元素不重复的整形数组排序

解法：遍历一遍，设置状态1，然后再次遍历，对状态等于1的进行输出，参考计数排序的原理。

（4）在2.5亿个整数中找出不重复的整数，注，内存不足以容纳这2.5亿个整数

解法1：采用2-Bitmap（每个数分配2bit，00表示不存在，01表示出现一次，10表示多次，11无意义）。

解法2：采用两个BitMap，即第一个Bitmap存储的是整数是否出现，接着，在之后的遍历先判断第一个BitMap里面是否出现过，如果出现就设置第二个BitMap对应的位置也为1，最后遍历BitMap，仅仅在一个BitMap中出现过的元素，就是不重复的整数。

解法3：分治+Hash取模，拆分成多个小文件，然后一个个文件读取，直到内存装的下，然后采用Hash+Count的方式判断即可。

该类问题的变形问题，如已知某个文件内包含一些电话号码，每个号码为8位数字，统计不同号码的个数。 8位最多99 999 999，大概需要99m个bit，大概10几m字节的内存即可。 （可以理解为从0-99 999 999的数字，每个数字对应一个Bit位，所以只需要99M个Bit==12MBytes，这样，就用了小小的12M左右的内存表示了所有的8位数的电话）

#### Reference

>https://www.infoq.cn/article/the-secret-of-bitmap/
>
>https://www.jianshu.com/p/6082a2f7df8e


