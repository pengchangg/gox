package main

import (
	"context"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

func main() {
	cache, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))

	go func(c *bigcache.BigCache) {
		i := 0
		for {
			i++
			c.Set(fmt.Sprintf("k_%d", i), []byte(fmt.Sprintf("v_%d", i)))
		}
	}(cache)

	go func(c *bigcache.BigCache) {
		for {
			time.Sleep(time.Nanosecond)
			fmt.Println(c.Len())
		}
	}(cache)
	select {}
}
