package dlock

import (
	"context"
	"time"

	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
)

type RedisPoolOption interface {
	Apply(*rpool)
}

type withPassword string

func (w withPassword) Apply(o *rpool) { o.host = string(w) }

// WithPassword provides an option to set a password to a Redis pool.
func WithPassword(v string) RedisPoolOption { return withPassword(v) }

type withTimeout time.Duration

func (w withTimeout) Apply(o *rpool) { o.timeout = time.Duration(w) }

// WithTimeout provides an option to set a timeout to a Redis pool.
func WithTimeout(v time.Duration) RedisPoolOption { return withTimeout(v) }

// RedisPool returns a connection pool for Redis.
func NewRedisPool(host string, opts ...RedisPoolOption) *redis.Pool {
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
	}
}

type rpool struct {
	host      string
	passwd    string
	timeout   time.Duration
	maxIdle   int
	maxActive int
	wait      bool
	idleTm    time.Duration
}

type RedisLockOption interface {
	Apply(*rlock)
}

type withHost string

func (w withHost) Apply(o *rlock) { o.hosts = append(o.hosts, string(w)) }

// WithHost provides an option to set a single Redis host for the lock.
func WithHost(v string) RedisLockOption { return withHost(v) }

type withHosts struct{ hosts []string }

func (w withHosts) Apply(o *rlock) { o.hosts = append(o.hosts, w.hosts...) }

// WithHosts provides an option to set a list of Redis hosts for the lock.
func WithHosts(v []string) RedisLockOption { return withHosts{v} }

type withPools struct{ pools []*redis.Pool }

func (w withPools) Apply(o *rlock) {
	for _, v := range w.pools {
		o.pools = append(o.pools, v)
	}
}

// WithPools provides an option to set a list of Redis pools for the lock.
func WithPools(v []*redis.Pool) RedisLockOption { return withPools{v} }

type withExtendAfter time.Duration

func (w withExtendAfter) Apply(o *rlock) { o.extend = time.Duration(w) }

// WithExtendAfter provides an option to set the duration before extending the lock.
func WithExtendAfter(v time.Duration) RedisLockOption { return withExtendAfter(v) }

func NewRedisLock(name string, mopts []redsync.Option, opts ...RedisLockOption) *rlock {
	lock := &rlock{
		hosts:  []string{},
		pools:  []redsync.Pool{},
		extend: time.Second * 5,
	}

	// Apply provided options, if any.
	for _, opt := range opts {
		opt.Apply(lock)
	}

	if len(lock.hosts) > 0 {
		for _, h := range lock.hosts {
			p := NewRedisPool(h)
			lock.pools = append(lock.pools, p)
		}
	}

	rs := redsync.New(lock.pools)
	lock.m = rs.NewMutex(name, mopts...)
	return lock
}

type rlock struct {
	hosts  []string
	pools  []redsync.Pool
	m      *redsync.Mutex
	extend time.Duration
	quit   context.Context
	cancel context.CancelFunc
}

func (l *rlock) Lock(ctx context.Context) error {
	if err := l.m.Lock(); err != nil {
		return err
	}

	l.quit, l.cancel = context.WithCancel(ctx)
	ticker := time.NewTicker(l.extend)

	// Continue lock until unlocked.
	go func() {
		for {
			select {
			case <-ticker.C:
			case <-l.quit.Done():
				return
			}

			// TODO: Do something with errors.
			// Propagate to Unlock?
			l.m.Extend()
		}
	}()

	return nil
}

func (l *rlock) Unlock() error {
	l.cancel()
	return nil
}
