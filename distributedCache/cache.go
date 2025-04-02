package distributedcache

import (
	"ProcessManger/utils/hashing"
	"sync"
)

var defaultCapacity = 10000

/*
what do i want in /this cache system
1. Any data type can be stored in value field allow pre determined data type as well as any data type
2. Cache should be thread safe add RW Mutex
3. We will implement a RR Cache replacement policy
4. A Cache system will have a capacity and a hash function
5. each instance will be distributed across multiple nodes
6. Will follow consistent hashing to distribute the cache across the nodes
7. will take the hash function in account while distributing the cache
8. will have a replication factor
9. will have a cache replacement policy
10. will have a cache eviction policy
11. will have a cache hit ratio
12. will have a cache miss ratio
13. will have a ttl
*/

type Cache struct {
	mu sync.Mutex
	items map[interface{}]any
	expiry int
	capacity int
}



