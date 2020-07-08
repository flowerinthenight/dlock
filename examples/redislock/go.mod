module github.com/flowerinthenight/redislock

go 1.14

require (
	github.com/flowerinthenight/dlock v0.0.0-20200708044148-5dc16ec8f18d
	github.com/gomodule/redigo v2.0.0+incompatible
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190918155943-95b840bb6a1f
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90
)
