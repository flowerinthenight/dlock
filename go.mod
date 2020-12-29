module github.com/flowerinthenight/dlock

go 1.14

require (
	cloud.google.com/go v0.74.0 // indirect
	cloud.google.com/go/spanner v1.12.0
	github.com/flowerinthenight/spindle v0.3.1
	github.com/go-redsync/redsync v1.4.2
	github.com/gomodule/redigo v2.0.0+incompatible
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b // indirect
	golang.org/x/sys v0.0.0-20201223074533-0d417f636930 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	golang.org/x/tools v0.0.0-20201229013931-929a8494cf60 // indirect
	google.golang.org/genproto v0.0.0-20201214200347-8c77b98c765d // indirect
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
