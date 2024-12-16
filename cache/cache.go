package cache

import (
	"errors"
	"time"
)

type Cache struct {
	val map[string]interface{}
}

func New() *Cache {
	return &Cache{
		make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}, t int) error {
	if key != "" {
		go func() {

			timer := time.NewTimer(time.Duration(t) * time.Second)
			<-timer.C
			delete(c.val, key)
		}()
		c.val[key] = value
		return nil
	}
	return errors.New("empty key")

}

func (c *Cache) Get(key string) (interface{}, bool) {
	value, exists := c.val[key]
	return value, exists
}

func (c *Cache) Delete(key string) error {
	_, exists := c.val[key]
	if exists {
		delete(c.val, key)
		return nil
	}
	return errors.New("invalid key")
}
