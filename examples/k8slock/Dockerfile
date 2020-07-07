FROM golang:latest
# COPY go.* /go/src/github.com/flowerinthenight/dlock/examples/k8slock/
COPY examples/k8slock/* /go/src/github.com/flowerinthenight/dlock/examples/k8slock/
WORKDIR /go/src/github.com/flowerinthenight/dlock/examples/k8slock/
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o k8slock .

FROM ubuntu:20.04
WORKDIR /k8slock/
COPY --from=0 /go/src/github.com/flowerinthenight/dlock/examples/k8slock .
ENTRYPOINT ["/k8slock/k8slock"]
