package cache

import (
	"encoding/json"
	"errors"
	"log"
	"sync"

	"github.com/avkim12/L0/postgres"
)

type Cache struct {
	sync.RWMutex
	items map[string]postgres.Order
}

func New() *Cache {

	items := make(map[string]postgres.Order)

	cache := Cache{
		items: items,
	}

	return &cache
}

func Backup(db *postgres.OrderDB, cache *Cache) {

	orders, err := db.GetAll()
	if err != nil {
		log.Println(err)
	}

	for _, value := range orders {
		cache.Set(value.UID, value.Model)
	}
}

func (c *Cache) Set(key string, value json.RawMessage) {

	c.Lock()
	defer c.Unlock()

	c.items[key] = postgres.Order{
		UID:   key,
		Model: value,
	}
}

func (c *Cache) Get(key string) (postgres.Order, bool) {

	c.RLock()
	defer c.RUnlock()

	order, found := c.items[key]
	if !found {
		return order, false
	}
	return order, true
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
