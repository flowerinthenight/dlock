# dlock
Package for distributed locks. At the moment, available implementations are [Redis](https://redis.io/topics/distlock) and Kubernetes using the [LeaseLock](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#lease-v1-coordination-k8s-io) resource.

A simple [`Locker`](https://github.com/flowerinthenight/dlock/blob/master/dlock.go) interface is also provided. All lock objects in this package implement this interface.

# Usage
### LeaseLock
A [sample code](https://github.com/flowerinthenight/dlock/tree/master/examples/k8slock) is provided for reference. A [deployment file](https://github.com/flowerinthenight/dlock/blob/master/examples/k8slock/k8slock.yaml) is also provided. It will deploy two pods that will both try to grab the same lock.

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

### TODO
PR's are welcome.
- [ ] Redis
- [ ] ectd
- [ ] Add CI
