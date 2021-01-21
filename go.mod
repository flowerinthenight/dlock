module github.com/flowerinthenight/dlock

go 1.14

require (
	cloud.google.com/go/spanner v1.13.0
	github.com/flowerinthenight/spindle v0.3.5
	github.com/go-redsync/redsync v1.4.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/uuid v1.1.5 // indirect
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	golang.org/x/oauth2 v0.0.0-20210113205817-d3ed898aa8a3 // indirect
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	golang.org/x/tools v0.1.0 // indirect
	google.golang.org/genproto v0.0.0-20210121164019-fc48d45331c7 // indirect
	google.golang.org/grpc v1.35.0 // indirect
	k8s.io/api v0.18.5 // indirect
	k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20200619165400-6e3d28b6ed19 // indirect
)

// curl -s https://proxy.golang.org/k8s.io/api/@v/kubernetes-1.16.0.info | jq -r .Version
// curl -s https://proxy.golang.org/k8s.io/apimachinery/@v/kubernetes-1.16.0.info | jq -r .Version
// curl -s https://proxy.golang.org/k8s.io/client-go/@v/kubernetes-1.16.0.info | jq -r .Version
replace (
	k8s.io/api => k8s.io/api v0.0.0-20190918155943-95b840bb6a1f
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190913080033-27d36303b655
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90
)
