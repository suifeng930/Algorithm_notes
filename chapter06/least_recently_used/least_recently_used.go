package main

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