package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	mu  sync.RWMutex
	val map[string]interface{}
}

func New() *Cache {
	return &Cache{
		val: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}, t int) error {

	if key == "" {
		return errors.New("empty key")
	}

	c.mu.Lock()
	c.val[key] = value
	c.mu.Unlock()

	go func() {
		time.Sleep(time.Duration(t) * time.Second)
		c.mu.Lock()
		delete(c.val, key)
		c.mu.Unlock()
	}()
	return nil
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.val[key]
	return value, exists
}

func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, exists := c.val[key]
	if exists {
		delete(c.val, key)
		return nil
	}
	return errors.New("invalid key")
}

func (c *Cache) SafeGet() map[string]interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	copy := make(map[string]interface{}, len(c.val))
	for keys, values := range c.val {
		copy[keys] = values
	}
	return copy
}
