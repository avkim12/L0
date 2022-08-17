package cache

import (
	"encoding/json"
	"errors"
	"sync"
)

type Cache struct {
	sync.RWMutex
	items map[string]Item
}

type Item struct {
	Value interface{}
}

func New() *Cache {
	items := make(map[string]Item)
	cache := Cache{
		items: items,
	}
	return &cache
}

func (c *Cache) Set(key string, value json.RawMessage) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = Item{
		Value: value,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	order, found := c.items[key]
	if !found {
		return nil, false
	}
	return order.Value, true
}

func (c *Cache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()

	if _, found := c.items[key]; !found {
		return errors.New("key not found")
	}
	delete(c.items, key)
	return nil
}
