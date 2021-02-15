# CNAA
ECGR 4090 Spring 2021 Cloud Native Application Architecture

### Lab 1
Work on readme -- basic  

### Lab 2
Implementing the Get() and Put() methods for the LRUCache based on interfaces.  

1.) Get() will check if the key exists in the cache, and append to the end of the queue (most recently used), and return the value of the key.  

2.) Put() will check if the key exists in the cache, and append to the end of the queue (most recently used), and update the cache value for the key. Otherwise, when the cache is full the least recently used (front of queue) will be evicted and the cache key with such index will be evicted, and the size of the cache will be decreased. Lastly, if the key does not exist in the cache (and the cache is not full) the key will be appended to the end of the queue (most recently used), the cache size will increase, and the value will be included to the key.

### Lab 3 
WIP
