![main](https://github.com/flowerinthenight/dlock/workflows/main/badge.svg)

# dlock
Package for distributed locks. At the moment, available implementations are [Kubernetes](https://kubernetes.io/) using the [LeaseLock](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#lease-v1-coordination-k8s-io) resource and [Redlock](https://redis.io/topics/distlock) via [redsync](https://github.com/go-redsync/redsync).

A simple [`Locker`](https://github.com/flowerinthenight/dlock/blob/master/dlock.go) interface is also provided. All lock objects in this package implement this interface.

# Usage
- ### LeaseLock
The simplest usage form looks something like:
```golang
lock := dlock.NewK8sLock("unique-id", "lock-name")
lock.Lock(context.TODO())
...
lock.Unlock
```

A [sample code](https://github.com/flowerinthenight/dlock/blob/master/examples/k8slock/main.go) is provided for reference. A [deployment file](https://github.com/flowerinthenight/dlock/blob/master/examples/k8slock/k8slock.yaml) is also provided (only tested on GKE). It will deploy two pods that will both try to grab the same lock.

```bash
# Deploy to k8s:
$ kubectl create -f k8slock.yaml

# See logs (not the full logs):
$ stern k8slock
main.go:53] [10.28.0.225] attempt to grab lock for a minute...
main.go:53] [10.28.4.52] attempt to grab lock for a minute...
main.go:47] [10.28.4.52] lock acquired by 10.28.4.52
main.go:57] [10.28.4.52] got the lock within that minute!
main.go:64] [10.28.4.52] now, let's attempt to grab the lock until termination
main.go:47] [10.28.4.52] lock acquired by 10.28.4.52
main.go:61] [10.28.0.225] we didn't get the lock within that minute
main.go:64] [10.28.0.225] now, let's attempt to grab the lock until termination
main.go:73] [10.28.0.225] stopping...
main.go:78] [10.28.0.225] we didn't get the lock in the end
main.go:73] [10.28.4.52] stopping...
main.go:75] [10.28.4.52] got the lock in the end

# Cleanup:
$ kubectl delete -f k8slock.yaml
```

- ### spindle
This implementation is a wrapper to the [spindle](https://github.com/flowerinthenight/spindle) distributed locking library by providing a blocking `Lock(...)` / `Unlock()` function pair. This is probably useful if you are already using [Cloud Spanner](https://cloud.google.com/spanner/).

The basic usage will look something like:
```golang
ctx := context.Background()
db, _ := spanner.NewClient(ctx, "your/database")
defer db.Close()

l := dlock.NewSpindleLock(&dlock.SpindleLockOptions{
  Client:   db,
  Table:    "testlease",
  Name:     "dlock",
  Duration: 1000,
})

start := time.Now()
l.Lock(ctx)
log.Printf("lock acquired after %v, do protected work...", time.Since(start))
time.Sleep(time.Second * 5)
l.Unlock()
```

- ### Redis
The Redis implementation is basically a wrapper to the brilliant [redsync](https://github.com/go-redsync/redsync) package, with additional utility functions for working with Redis connection pools. It's also implemented in a way to follow the [`Locker`](https://github.com/flowerinthenight/dlock/blob/master/dlock.go) interface.

To use with a single Redis host, no password, use defaults:
```golang
lock := dlock.NewRedisLock("testredislock", dlock.WithHost("1.2.3.4"))
lock.Lock(context.Background())
...
lock.Unlock()
```

To use with a single Redis host, with password, use defaults:
```golang
pool := dlock.NewRedisPool("1.2.3.4", dlock.WithPassword("secret-pass"))
lock := dlock.NewRedisLock("testredislock", dlock.WithPools([]*redis.Pool{pool}))
lock.Lock(context.Background())
...
lock.Unlock()
```

You can check out the [test file](https://github.com/flowerinthenight/dlock/blob/master/redislock_test.go) and the [sample code](https://github.com/flowerinthenight/dlock/blob/master/examples/redislock/main.go) for usage reference. You can run the sample code in separate terminals simultaneously and one process should be able to grab the Redis lock.

----

### TODO
PR's are welcome.
- [x] LeaseLock (k8s)
- [x] spindle
- [ ] etcd
- [ ] Zookeeper
- [ ] Consul
- [x] Redis
- [x] Add CI
