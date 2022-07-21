package cache

type Cache struct {
	info map[string]interface{}
}

func New() *Cache {
	return &Cache{
		make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.info[key] = value
}

func (c *Cache) Get(key string) interface{} {
	return c.info[key]
}

func (c *Cache) Delete(key string) {
	delete(c.info, key)
}
