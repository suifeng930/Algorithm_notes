## 漫画算法-小灰的算法之旅（01）

[TOC]



------

### 1. 什么是算法(algorithm)

------

在计算机领域里，算法是一系列程序指令，用于处理特定的运算和逻辑问题。衡量算法优劣的主要标准是**时间复杂度**和**空间复杂度**。

### 2. 什么是数据结构(data structure)

------

**数据结构**是数据的组织、管理和存储格式，其使用目的是为了高效的访问和修改数据。

**数据结构**包含数组、链表这样的**线性数据结构**，也包含树、图这样的复杂数据结构。

**数据结构(data structure)**是算法的基石。如果把算法比喻成美丽灵动的舞者，那么**数据结构**就是舞者脚下广阔而坚实的舞台。

#### 2.1 数据结构都有哪些组成方式？

------

##### 线性结构

线性结构是最简单的数据结构，包括数组、链表以及由它们衍生出来的栈、队列、哈希表

##### **树**

**树**是相对复杂的数据结构，其中比较有代表性的是二叉树，由它衍生出了**二叉堆**之类的数据结构。

##### 图

图是更为复杂的数据结构，因为在图中会呈现出多对多的关联关系。

##### 其他数据结构

除了基础数据结构之外，还有一些由基础数据结构变形而来用于解决某些特定问题，如跳表、哈希链表、位图等。例如在排序算法中的堆排序，利用的就是二叉堆这样一种数据结构；再如缓存淘汰算法LRU(least Recently Used),利用的就是特殊数据结构**哈希链表**

### 3. 什么是时间复杂度

------

**时间复杂度**是对一个算法运行时间长短的量度，用大O表示，记作T(n)=O(f(n))。常见的时间复杂度按照从低到高的顺序，包括O(1)、O(logn)、O(n)、O(nlogn)、O(n2)

#### 3.1 算法的好坏

衡量算法的好坏有很多标准，其中最重要的两个指标是算法的**时间复杂度**和**空间复杂度**。

#### 3.2 基本操作执行次数

由于受运行环境和输入规模的影响，代码的绝对执行时间是无法预告的。但我们可以预估代码的基本操作执行次数。

**基本操作执行次数**：设T(n)为程序基本操作执行次数的函数(也可认为是程序的相对执行时间函数)，n为输入规模。例如 T(n)=3n,执行次数是**线性**的；T(n)=5logn，执行次数是用**对数**计算的；T(n)=2,执行次数是**常量**；T(n)=0.5n2+0.5n,执行次数是用**多项式**计算的。

#### 3.3 渐进时间复杂度

**渐进时间复杂度（asymptotic time complexity）**: 若存在函数f(n),使得当n趋近于无穷大时，T(n)/f(n)的极限值为不等于零的常数，则称f(n)是T(n)的同数量级函数。记作T(n)=O(f(n)),称为O(f(n)); O为算法的渐进时间复杂度，简称时间复杂度。

直白的讲：时间复杂度就是把程序的相对执行时间函数T(n)简化为一个数量级，这个数量级可以是n,n2,n3.等。

#### 如何推导时间复杂度？

------

* 如果运行时间是常数量级，则用常数1表示
* 只保留时间函数中的最高项
* 如果最高阶项存在，则省去最高项前面的系数



### 4. 什么是空间复杂度

------

**空间复杂度**是对一个算法在运行过程中临时占用存储空间大小的度量，用大O表示，记作S(n)=O(f(n))。常见的空间复杂度按照从低到高的顺序，包括O(1)、O(n)、O(n2)等。其中递归算法的空间复杂度和递归深度成正比。

