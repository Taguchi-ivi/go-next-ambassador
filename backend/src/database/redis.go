package database

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var Cache *redis.Client
var CacheChannel chan string

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

}

func SetupCacheChannel() {
	CacheChannel = make(chan string)
	go func(ch chan string) {
		for {
			time.Sleep(5 * time.Second)

			Cache.Del(context.Background(), <-ch)
		}

	}(CacheChannel)

}

func ClearCache(key ...string) {
	for _, k := range key {
		CacheChannel <- k
	}
}
