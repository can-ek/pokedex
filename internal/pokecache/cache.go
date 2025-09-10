package pokecache

import (
	"sync"
	"time"
)

func NewCacheClient(expiry time.Duration) CacheClient {
	mux := &sync.Mutex{}
	newCache := &cache{expiry: expiry, mutex: mux, elements: map[string]cacheItem{}}

	ticker := time.NewTicker(expiry)
	go newCache.reapLoop(ticker)
	return newCache
}
