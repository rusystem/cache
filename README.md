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

    cache.Set("userId", 42)
    userId := cache.Get("userId")

    fmt.Println(userId)

    cache.Delete("userId")
    userId := cache.Get("userId")

    fmt.Println(userId)
}
````

``Set(key string, value interface{})`` - writing value to the cache by key

``Get(key string)`` - reading value from the cache by key

``Delete(key)`` - delete value from the cache by key
