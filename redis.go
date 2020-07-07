package dlock

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisPoolOption interface {
	Apply(*rpool)
}

type withPassword string

func (w withPassword) Apply(o *rpool)       { o.host = string(w) }
func WithPassword(v string) RedisPoolOption { return withPassword(v) }

type withTimeout time.Duration

func (w withTimeout) Apply(o *rpool)              { o.timeout = time.Duration(w) }
func WithTimeout(v time.Duration) RedisPoolOption { return withTimeout(v) }

type rpool struct {
	host      string
	passwd    string
	timeout   time.Duration
	maxIdle   int
	maxActive int
	wait      bool
	idleTm    time.Duration
}

// RedisPool returns a connection pool for Redis.
func NewRedisPool(host string, opts ...RedisPoolOption) (*redis.Pool, error) {
	if host == "" {
		return nil, fmt.Errorf("host empty, should be host:port")
	}

	pool := &rpool{
		host:      host,
		timeout:   5,
		maxIdle:   3,
		maxActive: 4,
		wait:      true,
		idleTm:    time.Second * 240,
	}

	// Apply provided options, if any.
	for _, opt := range opts {
		opt.Apply(pool)
	}

	var dialOpts []redis.DialOption
	if pool.passwd != "" {
		dialOpts = append(dialOpts, redis.DialPassword(pool.passwd))
	}

	dialOpts = append(dialOpts, redis.DialConnectTimeout(time.Second*pool.timeout))

	return &redis.Pool{
		MaxIdle:     pool.maxIdle,
		MaxActive:   pool.maxActive,
		Wait:        pool.wait,
		IdleTimeout: pool.idleTm,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", pool.host, dialOpts...)
		},
	}, nil
}
