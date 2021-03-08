module github.com/flowerinthenight/dlock

go 1.14

require (
	cloud.google.com/go/spanner v1.14.1
	github.com/flowerinthenight/spindle v0.3.6
	github.com/go-redsync/redsync v1.4.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	k8s.io/apimachinery v0.18.14
	k8s.io/client-go v0.18.14
	k8s.io/utils v0.0.0-20200619165400-6e3d28b6ed19 // indirect
)

// curl -s https://proxy.golang.org/k8s.io/api/@v/kubernetes-1.16.0.info | jq -r .Version
// curl -s https://proxy.golang.org/k8s.io/apimachinery/@v/kubernetes-1.16.0.info | jq -r .Version
// curl -s https://proxy.golang.org/k8s.io/client-go/@v/kubernetes-1.16.0.info | jq -r .Version
replace (
	k8s.io/api => k8s.io/api v0.18.14
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.14
	k8s.io/client-go => k8s.io/client-go v0.18.14
)
