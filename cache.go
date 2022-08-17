package cache

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type Cache struct {
	info map[string]interface{}
	mu   sync.RWMutex
}

func New() *Cache {
	return &Cache{
		info: make(map[string]interface{}),
	}
}

func (c *Cache) ttl(key string, ttl time.Duration) {
	for {
		select {
		case <-time.After(ttl):
			err := c.Delete(key)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}
}

func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.info[key]; !ok {
		return errors.New(fmt.Sprintf("Not found value by key=%s", key))
	}
	delete(c.info, key)

	return nil
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, ok := c.info[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Not found value by key=%s", key))
	}

	return c.info[key], nil
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.info[key] = value

	go c.ttl(key, ttl)
}
