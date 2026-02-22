package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	Data     map[string]cacheEntry
	Mu       sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	ticker := time.NewTicker(interval)

	cache := &Cache{
		Data:     map[string]cacheEntry{},
		Mu:       sync.Mutex{},
		interval: interval,
	}

	go func() {
		for t := range ticker.C {
			cache.reapLoop(t)
		}
	}()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	createdAt := time.Now()
	c.Mu.Lock()
	c.Data[key] = cacheEntry{
		createdAt: createdAt,
		val:       val,
	}
	c.Mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	val, ok := c.Data[key]
	if !ok {
		c.Mu.Unlock()
		return nil, false
	}
	c.Mu.Unlock()
	return val.val, true
}

func (c *Cache) reapLoop(now time.Time) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	for k, v := range c.Data {
		if now.Sub(v.createdAt) > c.interval {
			delete(c.Data, k)
		}
	}
}
