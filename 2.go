package main

import (
 "fmt"
 "sync"
)

type Cache struct {
 data map[string]string
 mu sync.RWMutex
}

func NewCache() *Cache {
 return &Cache{data: make(map[string]string)}
}

func (c *Cache) Get(key string) (string, bool) {
 c.mu.RLock()
 defer c.mu.RUnlock()
 val, ok := c.data[key]
 return val, ok
}

func (c *Cache) Set(key, value string) {
 c.mu.Lock()
 defer c.mu.Unlock()
 c.data[key] = value
}

func main() {
 cache := NewCache()
 cache.Set("user:1", "У")

 if val, ok := cache.Get("user:1"); ok {
  fmt.Println("Из кэша:", val)
 }
}
