# dlock
Package for distributed locks. At the moment, the available implementations are [Redis](https://redis.io/topics/distlock) and Kubernetes using the [LeaseLock](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#lease-v1-coordination-k8s-io) resource.

A simple [`Locker`](https://github.com/flowerinthenight/dlock/blob/master/dlock.go) interface is also provided. All lock objects in this package implement this interface.

# Usage
### LeaseLock
A [sample code](https://github.com/flowerinthenight/dlock/tree/master/examples/k8slock) is provided for reference. A [deployment file](https://github.com/flowerinthenight/dlock/blob/master/examples/k8slock/k8slock.yaml) is also provided.

```bash
# Deploy to k8s:
$ kubectl create -f k8slock.yaml

# See logs (not the full logs):
$ stern k8slock
I0707 22:00:59.444773       1 main.go:53] [10.28.0.225] attempt to grab lock for a minute...
I0707 22:00:59.399785       1 main.go:53] [10.28.4.52] attempt to grab lock for a minute...
I0707 22:01:35.322088       1 main.go:47] [10.28.4.52] lock acquired by 10.28.4.52
I0707 22:01:35.322110       1 main.go:57] [10.28.4.52] got the lock within that minute!
I0707 22:01:35.322156       1 main.go:64] [10.28.4.52] now, let's attempt to grab the lock until termination
I0707 22:01:40.329532       1 main.go:47] [10.28.4.52] lock acquired by 10.28.4.52
I0707 22:01:59.444828       1 main.go:61] [10.28.0.225] we didn't get the lock within that minute
I0707 22:01:59.444852       1 main.go:64] [10.28.0.225] now, let's attempt to grab the lock until termination
I0707 22:02:14.848754       1 main.go:73] [10.28.0.225] stopping...
I0707 22:02:14.848779       1 main.go:78] [10.28.0.225] we didn't get the lock in the end
I0707 22:02:14.859302       1 main.go:73] [10.28.4.52] stopping...
I0707 22:02:14.859346       1 main.go:75] [10.28.4.52] got the lock in the end

# Cleanup:
$ kubectl delete -f k8slock.yaml
```
