## 漫画算法-小灰的算法之旅（30）

[TOC]

## 1. 什么是LRU算法？

#### 什么是LRU算法

LRU(Least Recently Used),也就是最近最少使用的意思，是一种内存管理算法，该算法最早应用于Linux操作系统。这个算法基于一种假设：长期不被使用的数据，在未来被用到的几率也不大。因此，当数据所占内存达到一定阀值时，我们可以移除掉最近最少被使用的数据，以此来给出足够的内存空间。

#### LRU算法描述

LRU算法实际上是让你设计数据结构：首先要接受一个capacity参数作为缓存的最大容量，然后再实现两个API,一个是put(key,val)方法存入键值对，另一个是get(key)方法获取key对应的val.如果key不存在则返回-1。注意get()、put()方法的时间复杂度都必须是O(1)。要使put() get() 方法的时间复杂度为O(1)，那么对于该数据结构的要求就是：查找快，插入快，删除快，且有序。

**那么什么数据结构同时满足上述条件呢？**哈希表查找快，但数据无固定顺序；链表有顺序之分，插入、删除时间复杂度O(1),但查找慢。那要是把两者的优势结合起来呢？是不是就满足：查找快，插入快，删除快，且有序的特点呢？因此，将哈希表+链表结合起来，就演变成一种新的数据结构：**哈希链表**。

而LRU缓存算法的核心数据结构就是**哈希链表**，双向链表+哈希表的结合体。思想很简单，就是借助哈希表赋予了链表快速查找的特性：可以快速查找某个key是否存在缓存(链表)中，同时可以快速删除、添加新节点。

#### 示例演示

让我们以用户信息的需求为例，来演示一下LRU算法的基本思路。

1.假设使用哈希链表来缓存用户信息，目前缓存了4个用户，这4个用户是按照被访问的时间顺序依次从链表右端插入的。



2.如果这时业务方访问用户5，由于哈希链表中没有用户5的数据，需要从数据库中读取出来，插入到缓存中。此时，链表最右端是最新被访问的用户5，最左端是**最近最少被访问**的用户1.



3.接下来，如果业务方访问用户2，哈希链表中已经存在用户2的数据，这时我们把用户2从它的前驱节点和后继节点之间移除，重新插入链表的最右端。此时，链表的最右端变成了最新被访问的用户2，最左端仍然是**最近最少被访问**的用户1.



4.接下来，如果业务方请求修改用户4的信息。同样的道理，我们会把用户4从原来的位置移动到链表的最右侧，并把用户信息的值更新。这时，链表的最右端是最新被访问的用户4，最左端仍然是**最近最少被访问**的用户1.



5.后来业务方又要访问用户6，用户6在缓存中没有，需要插入哈希链表中。假设这时缓存容量已经达到上限，必须先删除**最近最少被访问**的数据，那么位于哈希链表最左端的用户1就会被删除，然后再把用户6插入到最右端的位置。



以上，就是LRU缓存算法的基本思路。

### 时间复杂度

LRU缓存算法的核心数据结构就是**哈希链表**，双向链表+哈希表的结合体。思想很简单，就是借助哈希表赋予了链表快速查找的特性：可以快速查找某个key是否存在缓存(链表)中，同时可以快速删除、添加新节点。因此时间复杂度为O(1)。

### 代码实现

```java

private Node head;
private Node end;
//缓存存储上限
private int limit;

private HashMap<String,Node> hashMap;

public LRUCache(int limit){
  this.limit=limit;
  hashMap =new HashMap<String,Node>();
}

public String get(String key){
  Node  node =hashMap.get(key);
  if(node==null){
    return null;
  }
  
  refreshNode(node);
  return node.value;
}

public void put(String key,String value){
  Node node=hashMap.get(key);
  
  if(node==null){
    //如果key不存在，则插入Key_value 
    if(hashMap.size()>=limit){
      String oldKey=removeNode(head);
      hashMap.remove(oldKey);
    }
    node =new Node(key,value);
    addNode(node);
    hashMap.put(key,node);
  }else{
    //如果key存在，则刷新key-value
    node.value=value;
    refreshNode(node);
  }
}

public void remove(String key){
  Node node=hashMap.get(key);
  removeNode(node);
  hashMap.remove(key);
}

/**
* 刷新被访问的节点位置
* param node 被访问的节点
*/
private void refreshNode(Node node){
  // 如果访问的是尾节点，则无须移动节点
  if(node==end){
    return;
  }
  // 移除节点
  removeNode(node);
  //重新插入节点
  addNode(node);
}

/**
* 删除节点
* param node 要删除的节点
*/
private String removeNode(Node node){
  if(node==head && node==end){
    //移除唯一的节点
    head=null;
    end=null;
    
  }else if(node==end){
    //移除尾节点
    end=end.pre;
    end.next=null;
  }else if(node==head){
    //移除头节点
    head=head.next;
    head.pre=null;
  }else{
    //移除中间节点
    node.pre.next=node.next; //node的前驱节点的next 指向 node 的next
    node.next.pre=node.pre; // node的next 前驱指向 node的前驱节点
  }
  return node.key;
}
/**
* 尾部插入节点
* param node 要插入的节点
*/
private void addNode(Node node){
  if(end!=null){
    end.next=node;
    node.pre=end;
    node.next=null;
  }
  end=node;
  if(head==null){
    head=node;
  }
}

class Node{
  Node(String key,String value){
    this.key=key;
    this.value=value;
  }
  public Node pre;
  public Node next;
  public String key;
  public String value;
}

public static void main(String[] args){
  
  LRUCache lruCache=new LRUCache(5);
  lruCache.put("001","用户1信息");
  lruCache.put("002","用户2信息");
  lruCache.put("003","用户3信息");
  lruCache.put("004","用户4信息");
  lruCache.put("005","用户5信息");
  lruCache.get("002");
  lruCache.put("004","用户4信息更新");
  lruCache.put("006","用户6信息");
  System.out.println(lruCache.get("001"));
  System.out.println(lruCache.get("006"));
}
```

#### golang实现

```go
type LinkNode struct {
	key,value int
	prev,next *LinkNode
}

type LRUCache struct {
	size int
	capacity int
	cache map[int]*LinkNode
	head ,end *LinkNode
}

func initLinkNode(key, value int) *LinkNode {
	return &LinkNode{
		key:   key,
		value: value,
	}
}

func NewLRUCacheOptions(capacity int) *LRUCache {

	cache :=LRUCache{
		capacity: capacity,
		cache:    map[int]*LinkNode{},
		head:     nil,
		end:      nil,
	}
	return &cache
}

func (l *LRUCache) Get(key int	)int  {
	if _, ok := l.cache[key]; !ok {
		return -1
	}
	node := l.cache[key]
	l.refreshNode(node)
	return node.value
}

func (l *LRUCache) Put(key,value int)  {
	node := l.cache[key]
	if node==nil {
		// 如果key不存在，则插入key-value
		if l.size>=l.capacity {
			oldKey := l.removeNode(l.head)
			delete(l.cache,oldKey)
			l.size--
		}
		node = initLinkNode(key, value)
		l.addNode(node) // 追加到链表中
		l.cache[key]=node //最佳到map中
		l.size++
	}else {
		// 如果key存在，则刷新key=value
		node:=l.cache[key]
		node.value=value
		l.refreshNode(node)
	}
}


//删除节点
func (l *LRUCache) removeNode(node *LinkNode) int {
	if node == l.head && node==l.end{
		//移除唯一的节点
		l.head=nil
		l.end=nil
	}else if node==l.end {
		//移除尾节点
		l.end=l.end.prev
		l.end.next=nil
	}else if node==l.head {
		//移除头节点
		l.head=l.head.next
		l.head.prev=nil
	}else {
		//移除中间节点
		node.prev.next=node.next
		node.next.prev=node.prev
	}
	return node.key
}

// 尾插法
func (l *LRUCache) addNode(node *LinkNode)  {
	if l.end!=nil {
		l.end.next=node
		node.prev=l.end
		node.next=nil
	}
	l.end=node
	if l.head==nil {
		l.head=node
	}
}

func (l *LRUCache) refreshNode(node *LinkNode)  {
	//移除节点
	l.removeNode(node)
	//重新添加节点
	l.addNode(node)
}
func main() {

	lruCache := NewLRUCacheOptions(5)
	lruCache.Put(1,1)
	lruCache.Put(2,2)
	lruCache.Put(3,3)
	lruCache.Put(4,4)
	lruCache.Put(5,5)
	lruCache.Get(2)
	lruCache.Put(4,124)
	lruCache.Put(6,126)
	log.Println(lruCache.Get(1))
	log.Println(lruCache.Get(6))
}
```


