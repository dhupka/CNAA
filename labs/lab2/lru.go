//Damian Hupka
//ECGR-4090: Cloud Native Application Architecture Spring 20201
//02/15/2021
//Lab #2 - Interfaces (Implementing get() and put() of LRUCache)
package lru

import ( 
	"errors"
)

type Cacher interface {
	Get(interface{}) (interface{}, error)
	Put(interface{}, interface{}) error
}

type lruCache struct {
	size      int
	remaining int
	cache     map[string]string
	queue     []string
}

func NewCache(size int) Cacher {
	return &lruCache{size: size, remaining: size, cache: make(map[string]string), queue: make([]string, size)}
}

func (lru *lruCache) Get(key interface{}) (interface{}, error) {
	// Your code here....
	
	//Check if key exists in the cache, if so appends to queue (most recently used)
	if node, found := lru.cache[key.(string)]; found {
		lru.queue = append(lru.queue, node)
		return lru.cache[key.(string)], nil					//Returning value of the key
	}
	return "-1", errors.New("Key does not exist.")

}

func (lru *lruCache) Put(key, val interface{}) error {
	// Your code here....
	//Check if key exists in cache, if so appends to queue (most recently used)
	if _, found := lru.cache[key.(string)]; found {
		lru.queue = append(lru.queue,lru.cache[key.(string)])
		lru.cache[key.(string)] = val.(string)				//Updating cache value if key exists
	} else {
		if lru.size == lru.remaining { 						//Evicting least recently used if cache is full (from queue and cache), reducing size of cache
			deleteIndex := lru.queue[0]
			lru.qDel(deleteIndex)							
			lru.size--										//Reducing size of cache as item removed
			delete(lru.cache,deleteIndex)
		}

	lru.queue = append(lru.queue,lru.cache[key.(string)]) 	//If key does not exist, appending it to queue (most recently used)
	lru.size++												//Increasing cache size as item added
	lru.cache[key.(string)] = val.(string)					//Updating cache value for non-existing key
	return nil	
	}
	return errors.New("ERROR.")
}


// Delete element from queue
func (lru *lruCache) qDel(ele string) {
	for i := 0; i < len(lru.queue); i++ {
		if lru.queue[i] == ele {
			oldlen := len(lru.queue)
			copy(lru.queue[i:], lru.queue[i+1:])
			lru.queue = lru.queue[:oldlen-1]
			break
		}
	}
}
