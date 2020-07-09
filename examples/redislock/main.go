package main

import (
	"context"
	"log"
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
	l := dlock.NewRedisLock("testredislock", dlock.WithPools([]*redis.Pool{pool}))

	// Use the usual lock/unlock.
	l.Lock(context.TODO())
	log.Println("locked")
	time.Sleep(time.Second * 5)
	l.Unlock()
	log.Println("unlocked")

	// Use context expiration.
	// ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	// log.Println("lock with context, expire 10s")
	// l.Lock(ctx)
	// <-ctx.Done()
	// log.Println("unlocked")
}
