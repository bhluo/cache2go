package cache2go

import (
	"log"
	"os"
	"sync"
	"time"
)

var (
	//表名为string
	cache = make(map[string]*CacheTable)
	mutex sync.RWMutex
)

// Cache returns the existing cache table with given name or creates a new one
// if the table does not exist yet.
func Cache(table string) *CacheTable {
	mutex.RLock()
	t, ok := cache[table]
	mutex.RUnlock()

	//如果不存在name为"table"的CacheTable
	if !ok {
		mutex.Lock()
		t, ok = cache[table]
		// Double check whether the table exists or not.
		if !ok {
			t = &CacheTable{
				name:            table,
				items:           make(map[interface{}]*CacheItem),
				cleanupInterval: 5 * time.Minute, // 默认清理间隔为5分钟
				logger:          log.New(os.Stderr, "cache2go: ", log.LstdFlags),
			}
			cache[table] = t

			go t.expirationCheck()
		}
		mutex.Unlock()
	}

	return t
}
