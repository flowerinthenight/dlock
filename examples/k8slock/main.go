package main

import (
	"context"
	goflag "flag"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/flowerinthenight/dlock"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"k8s.io/klog/v2"
)

func init() {
	klog.InitFlags(nil)
	goflag.Parse()
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
}

var (
	rootCmd = &cobra.Command{
		Use:   "k8slock",
		Short: "example of k8slock",
		Long:  "Example of k8slock.",
		RunE:  root,
	}
)

func run(ctx context.Context, done chan error) {
	id := os.Getenv("MY_POD_IP") // set from k8s deployment
	name := "k8slock"            // the name of the LeaseLock resource
	var locked int32

	// Create the locker object with some callback functions. By default,
	// this uses the 'default' namespace.
	lock := dlock.NewK8sLock(id, name,
		dlock.WithLeaseDuration(time.Second*30),
		dlock.WithStartCallback(func(ctx context.Context) {
			atomic.StoreInt32(&locked, 1)
		}),
		dlock.WithNewLeaderCallback(func(identity string) {
			if identity == id {
				klog.Infof("[%v] lock acquired by %v", id, identity)
			}
		}),
	)

	tm, _ := context.WithTimeout(context.TODO(), time.Minute)
	klog.Infof("[%v] attempt to grab lock for a minute...", id)
	lock.Lock(tm) // block for 1min, trying to grab lock

	if atomic.LoadInt32(&locked) == 1 {
		klog.Infof("[%v] got the lock within that minute!", id)
		atomic.StoreInt32(&locked, 0) // reset
		lock.Unlock()
	} else {
		klog.Infof("[%v] we didn't get the lock within that minute", id)
	}

	klog.Infof("[%v] now, let's attempt to grab the lock until termination", id)
	time.Sleep(time.Second * 5)

	go lock.Lock(context.TODO())

	for {
		select {
		case <-ctx.Done():
			klog.Infof("[%v] stopping...", id)
			if atomic.LoadInt32(&locked) == 1 {
				klog.Infof("[%v] got the lock in the end", id)
				lock.Unlock()
			} else {
				klog.Infof("[%v] we didn't get the lock in the end", id)
			}

			done <- nil
			return
		}
	}
}

func root(cmd *cobra.Command, args []string) error {
	quit, cancel := context.WithCancel(context.TODO())
	done := make(chan error)
	go run(quit, done)

	// Wait for termination.
	go func() {
		sigch := make(chan os.Signal)
		signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
		<-sigch
		cancel()
	}()

	return <-done
}

func main() {
	rootCmd.Execute()
}
