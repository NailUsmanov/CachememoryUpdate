package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	mu  sync.Mutex
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
		timer := time.NewTimer(time.Duration(t) * time.Second)
		<-timer.C
		c.mu.Lock()
		defer c.mu.Unlock()
		delete(c.val, key)
	}()
	return nil
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
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
