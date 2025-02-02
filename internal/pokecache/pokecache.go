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
	Entries  map[string]cacheEntry
	Mu       sync.RWMutex
	Interval time.Duration
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	entry, exists := c.Entries[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.Interval)
	for range ticker.C {
		now := time.Now()
		c.Mu.Lock()
		for key, entry := range c.Entries {
			if now.Sub(entry.createdAt) > c.Interval {
				delete(c.Entries, key)
			}
		}
		c.Mu.Unlock()
	}
}

func NewCache(d time.Duration) *Cache {
	c := &Cache{
		Entries:  make(map[string]cacheEntry),
		Interval: d,
	}
	go c.reapLoop()
	return c
}
