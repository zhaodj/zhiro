package zhiro

import (
	"sync"
)

type Cache interface {
	Get(key interface{}) interface{}
	Put(key interface{}, val interface{}) interface{}
	Remove(key interface{}) interface{}
	Clear()
	Size() int
}

type MemoryCache struct {
	items map[interface{}]interface{}
	lock  sync.RWMutex
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{items: make(map[interface{}]interface{})}
}

func (c *MemoryCache) Get(key interface{}) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if item, ok := c.items[key]; ok {
		return item
	}
	return nil
}

func (c *MemoryCache) Put(key interface{}, val interface{}) interface{} {
	c.lock.Lock()
	defer c.lock.Unlock()

	item, ok := c.items[key]
	if !ok {
		item = nil
	}

	c.items[key] = val
	return item
}

func (c *MemoryCache) Remove(key interface{}) interface{} {
	c.lock.Lock()
	defer c.lock.Unlock()

	item, ok := c.items[key]
	if !ok {
		item = nil
	}
	delete(c.items, key)
	return item
}

func (c *MemoryCache) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.items = map[interface{}]interface{}{}
}
