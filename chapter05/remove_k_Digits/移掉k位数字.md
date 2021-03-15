## 漫画算法-小灰的算法之旅（25）

[TOC]

### 1. 删去k个数字后的最小值

>Q: 给出一个整数，从该整数中去掉k个数字，要求剩下的数字形成的新整数尽可能小。
>
>其中整数的长度大于或等于k,给出的整数的大小可以超过long类型的数字范围。
>
>示例1:
>
>给定整数为1593212 ,删去3个数字，新整数最小的情况为1212
>
>给定整数为30200，删去1个数字，新整数最小的情况是200
>
>给定整数为10，删去2个数字(注意这里删除的是2个数字)，新整数的最小情况是0
>
>给定整数为1432219，删去3个数字，新整数的最小情况是1219

#### 方法一(贪心算法)

##### 思路

对于两个相同长度的数字序列，最左边不同的数字决定了这两个数字的大小，例如,对于A=1axxx,B=1bxxx,如果a>b则A>B。基于此，我们可以知道，若要使得剩下的数字最小，需要保证靠前的数字尽可能小。

**举个例子**：

给出一个整数541270936,要求删去1个数字，让剩下的整数尽可能小。

此时，无论删除哪一个数字，最后的结果都是从9位整数变成8位整数。既然同样是8位整数，显然应该有限吧高位的数字降低，这样对新整数的值影响最大。



如何把高位的数字降低呢？很简单，**把原整数的所有数字从左到右进行比较，如果发现某一位数字大于它右面的数字，那么在删除该数字后，必然会使该数位的值降低。**因为右面比它小的数字顶替了它的位置。在上面的例子中，数字5右侧的数字4<5，所以删除数字5，最高位数字降低成了4。



对于整数541 270 936 ,删除一个数字所能得到的最小值是41 270 936 。那么对于41 270 936 ,删除一个数字的最小值，是多少呢?

从左到右遍历，数字4是第1个比右侧数字大的数(4>1).



对于整数1 270 936 中再删除一个数字，得到中最小值是多少呢？ 因为1<2 ,2<7，7>0 所以被删除的数字因为7.



因此原数541 270 936 ，删除3个数字后的最小值，新数字的最小值为 120 936



像这样依次求得**局部最优解**，最终得到**全局最优解**的思想，叫作**贪心算法**。

##### 时间复杂度

因为每一次循环都需要从头开始遍历所有的数字，需要经历k轮循环。因此时间复杂度为O(kn)。

##### 代码实现

```java

/**
* 删除整数的k个数字，获得删除后的最小值
* param num 原整数
* param k 删除数量
*/
public static String removeKDigits(String num,int k){
  for(int i=0;i<k;i++){
    boolean hasCut =false;
    //从左向右遍历，找到比自己右侧数字大的数字并删除
    for(int j=0;j<num.length()-1;j++){
      if(num.charAt(j)>num.charAt(j+1)){
        num=num.substring(0,num)+num.substring(j+1,num.length());
        hasCut=true;
        break;
      }
    }
    // 如果没有找到要删除的数字，则删除最后一个数字
    if(!hasCut){
      num=num.substring(0,num.length()-1);
    }
  }
  
  //清除整数左侧的数字0
  int start =0;
  for(int j=0;j<num.length()-1;j++){
    if(num.charAt(j)!='0'){
      break;
    }
    start++;
  }
  num =num.substring(start.num.length());
  //如果整数的所有数组都被删除了，直接返回0
  if(num.length()==0){
    return "0";
  }
  return num;
}

public static void main(String[] args){
  System.out.println(removeKDigits("1593212",3));
  System.out.println(removeKDigits("30200",1));
  System.out.println(removeKDigits("10",2));
  System.out.println(removeKDigits("541270936",3));
}
```

#### 方法二(贪心算法+单调栈)

##### 思路

继续沿用方法一的思想，不过这里我们要做一些优化，方法一采用的**是以k作为外循环，遍历数字作为内循环，需要额外考虑的东西很多**。我们可以换一个思路，以遍历数字作为外循环，以k最为内循环，并引入栈来辅助我们处理。

因此，对于每个数字，如果该数字小于栈顶元素，我们就不断地弹出栈顶元素，直到

- 栈为空
- 或者新的栈顶元素不大于当前数字
- 或者我们已经删除了 k*k* 位数字

此外，我们还需要注意以下几个点：

* 如果我们删除了k个数字且k<n，这种情况下我们需要从序列尾部删除额外的n-k个数字
* 如果最终的数字序列存在前导零,我们需要删除前导零
* 如果最终数字序列为空，我们应该返回"0"
* 此外还需要注意将数字从栈中依次弹出并进行翻转

##### 时间复杂度

这样一来，我们只会对所有数字遍历一次，遍历的时间复杂度为O(n),把栈转化为字符串的时间复杂度也是O(n),所以最终的时间复杂度为O(n)。

##### 代码实现

```java

/**
* 删除整数的k个数字，获得删除后的最小值
* param num 原整数
* param k 删除数量
*/
public static String removeKDigits(String num,int k){
  
  // 新整数的最终长度，=原整数长度-k
  int nemLength=num.length()-k;
  // 创建一个栈，用于接收所有的数字
  char[] stack=new char[num.length()];
  int top=0;
  for(int i=0;i<num.length();++i){
    //遍历当前数字
    char c =num.charAt(i);
    //当栈顶数字大于遍历到的当前数字时，栈顶数字出栈(相当于删除数字)
    while(top>0 && stack[top-1]>c && k>0){
      top-=1;
      k-=1;
    }
    //遍历到的当前数字入栈
    stack[top++]=c;
  }
  // 找到栈中第1个非零数字的位置，以此构建新的整数字符串
  int offset =0;
  while(offset <newLength && stack[offset] =='0'){
    offset++;
  }
  return offset==newLength ? "0":new String(stack,offset,newLength-offset);
}

public static void main(String[] args){
  System.out.println(removeKDigits("1593212",3));
  System.out.println(removeKDigits("30200",1));
  System.out.println(removeKDigits("10",2));
  System.out.println(removeKDigits("541270936",3));
}
```

##### golang实现

```go

func removeKDigits(num string, k int) string {

	//初始化一个栈，用于接收所有的数字
	stack :=[]byte{}
	//遍历原数列
	for i := range num {
		//遍历当前数字
		digit :=num[i]
		// 如果k>0 且 栈顶数字大于遍历到的当前数字，
		for k>0 && len(stack)>0 && digit<stack[len(stack)-1] {
			// 栈顶数字出栈
			stack=stack[:len(stack)-1]
			// k数字减1
			k--
		}
		// 遍历到的当前数字入栈
		stack=append(stack,digit)
	}
	// 获取截取之后的栈
	stack=stack[:len(stack)-k]

	// 删除数字中包含前导为0的数字
	digits :=strings.TrimLeft(string(stack),"0")
	// 如果最终的字符串为 空 返回 0
	if digits=="" {
		digits="0"
	}
	return digits
}

func main() {
	log.Println(removeKDigits("1593212",3))
	log.Println(removeKDigits("30200",1))
	log.Println(removeKDigits("10",2))
	log.Println(removeKDigits("541270936",3))
}

```


