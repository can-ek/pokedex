package pokecache

import (
	"sync"
	"time"
)

type cacheItem struct {
	createdAt time.Time
	expireAt  time.Time
	val       []byte
}

type cache struct {
	elements map[string]cacheItem
	expiry   time.Duration
	mutex    *sync.Mutex
}

type CacheClient interface {
	Add(key string, val []byte) error
	Get(key string) ([]byte, bool)
	reapLoop(ticker *time.Ticker)
}

func (c *cache) Add(key string, val []byte) error {
	c.mutex.Lock()
	c.elements[key] = cacheItem{val: val, createdAt: time.Now(), expireAt: time.Now().Add(c.expiry)}
	//	fmt.Println("Added to cache:", key)
	c.mutex.Unlock()
	return nil
}

func (c *cache) Get(key string) ([]byte, bool) {
	if item, ok := c.elements[key]; ok {
		// fmt.Println("Found in cache:", key)
		return item.val, true
	}

	return nil, false
}

func (c *cache) reapLoop(ticker *time.Ticker) {
	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.elements {
			if time.Now().After(entry.expireAt) {
				delete(c.elements, key)
				// fmt.Println("Deleted from cache:", key)
			}
		}
		c.mutex.Unlock() // Ensure it gets unlocked
	}
}
