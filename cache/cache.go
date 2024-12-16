package cache

import "errors"

type Cache struct {
    val map[string]interface{}
}

func New() *Cache {
    return &Cache{
        make(map[string]interface{}),
    }
}

func (c *Cache) Set(key string, value interface{}) error {
    if key != "" {
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
