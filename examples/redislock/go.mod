module github.com/flowerinthenight/dlock/examples/redislock

go 1.14

require (
	github.com/flowerinthenight/dlock v0.2.0
	github.com/gomodule/redigo v2.0.0+incompatible
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190918155943-95b840bb6a1f
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90
)
