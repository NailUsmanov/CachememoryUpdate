package main

import (
	// Импортируем пакет из папки cache

	"cache/cache"
	"fmt"
	"time"

	"math/rand"
)

func main() {
	newCache := cache.New()
	methods := []func(){
		func() {
			key := fmt.Sprintf("key %d", rand.Intn(100))
			value := rand.Intn(10000)
			err := newCache.Set(key, value, 5)
			if err != nil {
				fmt.Println("Set Error:", err)
			} else {
				fmt.Printf("Set: {Key: %s, Value: %d}\n", key, value)
			}
		},
		func() {
			key := fmt.Sprintf("key%d", rand.Intn(100))
			value, found := newCache.Get(key)
			if found {
				fmt.Printf("Get: {Key: %s, Value: %v}\n", key, value)
			} else {
				fmt.Printf("Get: Key %s not found\n", key)
			}

		},
		func() {
			key := fmt.Sprintf("key%d", rand.Intn(100))
			err := newCache.Delete(key)
			if err != nil {
				fmt.Println("Delete Error:", err)
			} else {
				fmt.Printf("Delete: {Key: %s}\n", key)
			}
		},
	}

	for i := 0; i < 10; i++ {
		method := methods[rand.Intn(len(methods))]
		method() // Вызов случайной функции
	}

	newCache.Set("userId1", "123", 3)
	newCache.Set("userId2", "456", 5)
	newCache.Set("userId3", "789", 5)

	fmt.Println(newCache.Get("userId1"))
	time.Sleep(2 * time.Second)
	fmt.Println(newCache.Get("userId1"))

	err := newCache.Delete("userId2")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Cache state:", newCache.SafeGet())

}
