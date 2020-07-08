package dlock

import (
	"os"
	"testing"
)

func Test__NewRedisPool(t *testing.T) {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		return
	}

	pool := NewRedisPool(host, WithPassword(os.Getenv("REDIS_PASSWORD")))
	con := pool.Get()
	if con == nil {
		t.Fatal("got nil con")
	}

	con.Close()
}

func Test__NewRedisLock(t *testing.T) {
	l := NewRedisLock("hello", nil)
	if l == nil {
		t.Fatal("got nil lock")
	}
}
