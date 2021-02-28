## 漫画算法-小灰的算法之旅（02）

[TOC]



------

### 1. 什么是数组

------

**数组(array)**:是**有限**个**相同类型**的变量所组成的**有序集合**，数组中的每一个变量被称为**元素**。数组是最为简单、最为常用的数据结构。

**数组的特点**：在内存中**顺序存储**，因此可以很好的地实现逻辑上的**顺序表**。

**内存**是由一个个连续的内存单元组成的，每一个**内存单元**都有自己的地址。数组中的每一个**元素**，都对应着一个个内存单元中；并且元素之间紧密联系排列，既不能打乱元素的存储顺序，也不能跳过某个存储单元进行存储。

### 2. 数组的基本操作

------

数据结构的基本操作无非就是增、删、查、改四种情况。

#### 读取元素

对于数组来说，读取元素是最简单的操作。因为数组在内存中顺序存储，所以只要给出一个数组下标，就可以读取到对应的数组元素。根据下标读取元素的方式叫做**随机读取**。

```java
int [] array =new int[]{1,2,3,4,5,8,9};
// 输出数组下标为3的元素
System.out.println(array[3]);
```

#### 更新元素

要把数组中的某一个元素的值替换为一个新值，也是比较简单的操作。直接利用数组下标，就可以把新值复制给该元素了。

```java
int [] array =new int[]{1,2,3,4,5,8,9};
//给下标为3的元素赋值
array[3]=9
// 输出数组下标为3的元素
System.out.println(array[3]);
```

**数组读取元素和更新元素的时间复杂度都是O(1)**

#### 插入元素

插入数组元素的操作存在三种情况：

* 尾部插入
* 中间插入
* 超范围插入

**尾部插入**：是最简单的情况，直接把插入的元素放在数组的尾部空闲位置即可；等同于更新元素的操作。

**中间插入**：由于数组的每一个元素都有其固定下标，所以不得不首先把插入位置及后面的元素向后移动，腾出地方，再把要插入的元素放到对应的数组位置上。

```java
// 中间插入操作的完整代码实现
private int[] array;
private int size;

public MyArray(int capacity){
  this.array=new int[capacity];
  size=0;
}

/**
* 数组插入元素
* param index 插入的位置
* param element 插入的元素
*/
public void insert(int index,int element) throws Exception{
  // 判断访问下标是否超出范围
  if (index<0 || index>size){
    throw new IndexOutOfBoundsException("超出数组实际元素范围")；
  }
  // 从有向左循环，将元素逐个向右挪一位
  for (int i=size-1; i>=index;i--){
    array[i+1]=array[i];
  }
  //腾出的位置放入新元素
  array[index]=element;
  size++;
}

/** 
* 输出数组
*/
public void output(){
  for(int i=0;i<size;i++){
    System.out.println(array[i]);
  }
}

public static void main(String[] args) throws Exception{
  MyArray myArray=new MyArray(10);
  myArray.insert(0,3);
  myArray.insert(1,7);
  myArray.insert(2,9);
  myArray.insert(3,5);
  myArray.insert(1,6);
  myArray.output();
  
}
// 代码中的成员变量size 是数组实际元素的数量，如果插入的元素在数组尾部，传入的下标参数index等于size;如果插入元素在数组中间或头部，则index小于size。如果传入的下标参数大于size或小于0，则认为是非法输入，会直接跑出异常。
```

**超范围插入**：假如存在一个长度为6 的数组，已经装满的元素，此时还想插入一个新元素。就涉及到数组的扩容了；此时可以创建一个新数组，长度为旧数组的2倍，再把旧数组中的元素统统复制到新数组，这样就实现了数组的扩容操作。

```java
//数组扩容代码
private int[] array;
private int size; // 数组实际元素的个数

public MyArray(int capacity){
  this.array=new int[capacity];
  size=0;
}

/**
* 数组插入元素
* param index 插入的位置
* param element 插入的元素
*/
public void insert(int index,int element) throws Exception{
  
  // 判断访问下标是否超出范围
  if(index<0 || index>size){
    throw new IndexOutOfBoundsException("超出数组实际元素范围")；
  }
  
  // 如果实际元素达到数组容量上限，则对数组进行扩容
  if(size>=array.length){
    resize();
  }
  // 从右向左循环，将元素逐个向右挪动1位
  for(int i=size-1;i>index;i--){
    array[i+1]=array[i];
  }
  //腾出的位置放入新插入的元素
  array[index]=element;
  size++;
}

/** 
* 数组扩容
*/
public void resize(){
  int[] arrayNew=new int[array.length*2];
  //从旧数组复制到新数组
  System.arraycopy(array,0,arrayNew,0,array.length);
  array=arrayNew;
}

/** 
* 输出数组
*/
public void output(){
  for(int i=0;i<size;i++){
    System.out.println(array[i]);
  }
}
public static void main(String[] args) throws Exception{
  MyArray myArray=new MyArray(10);
  myArray.insert(0,3);
  myArray.insert(1,7);
  myArray.insert(2,9);
  myArray.insert(3,5);
  myArray.insert(1,6);
  myArray.output();
  
}

```

#### 删除元素

数组的删除操作和插入操作的过程相反，如果删除的元素位于数组中间，其后的元素都需要向前挪动一位。

```java
// 由于不涉及扩容问题，所以删除操作的代码实现比插入操作要简单。

/**
* 数组删除元素
* param index 删除元素的位置
*/
public int delete(int index) throws Exception{
  // 判断访问下标是否超出范围
  if(index<0 || index>size){
    throw new IndexOutOfBoundsException("超出数组实际元素范围")；
  }
  int deletedElement =array[index];
  // 从左向右循环，将元素逐个向左挪动一位
  for(int i=index;i<size-1;i++){
    array[i]=array[i+1];
  }
  size--;
  return deletedElement;
}
```



**数组的插入和删除操作，时间复杂度是多少？**

**插入操作**：数组扩容的时间复杂度是O(n);插入并移动元素的时间复杂度也是O(n);综合起来插入时间复杂度是O(n);

**删除操作**：只涉及元素的移动，时间复杂度也是O(n)。 对于删除操作，还有一种技巧，前提是数组元素没有顺序要求。对于需要删除的是数组中index下标的数据，可以把**size-1**下标的数据复制到元素下标为index所在的位置，然后再删除掉最后一个元素；此方式无需进行大量数组元素的移动操作，因此时间复杂度为O(1).

### 3. 数组的优势和劣势

------

**优势**：数组拥有非常高效的随机访问能力，只要给出下标，就可以用常量时间找到指定元素。有一种高效查找元素的算法叫作**二分查找**，就是利用数组的这个优势。

**劣势**：数组的劣势，体现在插入和删除元素方面。因此数组元素是连续紧密的存储在内存中，插入、删除元素都会导致大量元素被迫移动，影响效率。

**总结**：数组适合的是**读多、写少**的场景。


