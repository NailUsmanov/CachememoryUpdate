package main

import (
	// Импортируем пакет из папки cache

	"cache/cache"
	"fmt"
)

func main() {
	newCache := cache.New()

	newCache.Set("userId1", 123)
	newCache.Set("userId2", "456")
	newCache.Set("userId3", "789")

	value, exists := newCache.Get("userId3")
	if exists {
		fmt.Println("Value for userId3:", value)
	} else {
		fmt.Println("Key userId3 not found")
	}

	err := newCache.Delete("userId1")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Cache state:", newCache)
}
