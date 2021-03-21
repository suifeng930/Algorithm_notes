package main

import "log"

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
