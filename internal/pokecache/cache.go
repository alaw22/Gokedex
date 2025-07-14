package pokecache

import (
	"sync"
	"time"
	// "fmt"
)


type cacheEntry struct {
	val 		[]byte
	createdAt 	time.Time // Represents an instant in time with nanosecond precision
}

type Cache struct {
	Mu 		*sync.Mutex
	Memory 	map[string]cacheEntry
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		val: val,
		createdAt: time.Now(),
	}
	
	c.Mu.Lock()
	c.Memory[key] = entry
	c.Mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	entry, ok := c.Memory[key]
	c.Mu.Unlock()
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	
	for {
		t := <-ticker.C // blocks for every interval
		c.Mu.Lock()
		for key, entry := range c.Memory {
			if t.Sub(entry.createdAt) > interval {
				delete(c.Memory,key)
			}
		}
		c.Mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Mu: &sync.Mutex{},
		Memory: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}