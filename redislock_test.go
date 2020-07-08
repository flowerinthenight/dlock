package dlock

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	host = os.Getenv("REDIS_HOST")
	pass = os.Getenv("REDIS_PASSWORD")
)

func Test__LockUnlock(t *testing.T) {
	if host == "" {
		return
	}

	// Test with a single Redis instance with password.
	pool := NewRedisPool(host, WithPassword(pass))

	// Use 1 Redis pool.
	l := NewRedisLock("testredislock", nil, WithPools([]*redis.Pool{pool}))
	ch := make(chan error, 1)

	go func() {
		time.Sleep(time.Second)
		l.Lock(context.Background())
		t.Log("locked from goroutine")
		time.Sleep(time.Second * 3)
		l.Unlock()
		t.Log("unlocked from goroutine")
		ch <- nil
	}()

	// This should lock first.
	err := l.Lock(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	t.Log("locked from main")
	l.Unlock()
	t.Log("unlocked from main")
	<-ch
}
