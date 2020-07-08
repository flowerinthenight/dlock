package main

import (
	"context"
	"os"
	"time"

	"github.com/flowerinthenight/dlock"
	"github.com/gomodule/redigo/redis"
)

func main() {
	host := os.Getenv("REDIS_HOST")
	pass := os.Getenv("REDIS_PASSWORD")
	if host == "" {
		return
	}

	// Test with a single Redis instance with password.
	pool := dlock.NewRedisPool(host, dlock.WithPassword(pass))

	// Use 1 Redis pool.
	l := dlock.NewRedisLock("testredislock", nil, dlock.WithPools([]*redis.Pool{pool}))

	l.Lock(context.TODO())
	time.Sleep(time.Second * 5)
	l.Unlock()
}
