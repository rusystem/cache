# Getting cache
``` 
go get github.com/rusystem/cache
```

## How to use. Example:

````
package main

import "fmt"

func main() {
    cache := cache.New()

    cache.Set("userId", 42, time.Second * 5)
    userId := cache.Get("userId")

    fmt.Println(userId)

    cache.Delete("userId")
    userId := cache.Get("userId")

    fmt.Println(userId)
}
````

``Set(key string, value interface{}, ttl time.Duration)`` - writing value to the cache by key, with time to live

``Get(key string)`` - reading value from the cache by key

``Delete(key)`` - delete value from the cache by key
