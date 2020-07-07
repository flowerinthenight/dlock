module github.com/flowerinthenight/dlock

go 1.14

require (
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/gomodule/redigo v1.8.2
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
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
