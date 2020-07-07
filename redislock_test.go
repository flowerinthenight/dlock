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

	pool, err := NewRedisPool(host, WithPassword(os.Getenv("REDIS_PASSWORD")))
	if err != nil {
		t.Fatal(err)
	}

	con := pool.Get()
	if con == nil {
		t.Fatal("got nil con")
	}

	con.Close()
}
