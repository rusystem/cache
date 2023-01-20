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

    if err := cache.Set("userId", 42, 5); err != nil { // 5 - time in seconds
        log.Fatal(err)
    }
    
    userId, err := cache.Get("userId")
    if err == nil {
        i := userId.(int)
    }

    fmt.Println(i)

    if err = cache.Delete("userId"); err != nil {
        log.Fatal(err)
    }
}
````

``Set(key interface{}, value interface{}, ttl int64)`` - writing value to the cache by key, with time to live

``Get(key string)`` - reading value from the cache by key

``Delete(key)`` - delete value from the cache by key
