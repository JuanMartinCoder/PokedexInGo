package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	dataCache map[string]cacheEntry
	interval  time.Duration
	mutex     sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		dataCache: make(map[string]cacheEntry),
		interval:  interval,
		mutex:     sync.Mutex{},
	}

	go cache.readLoop(cache.interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.dataCache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, ok := c.dataCache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) readLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)

	c.mutex.Lock()
	for k, v := range c.dataCache {
		if v.createdAt.Before(timeAgo) {
			delete(c.dataCache, k)
		}
	}
	c.mutex.Unlock()
}
