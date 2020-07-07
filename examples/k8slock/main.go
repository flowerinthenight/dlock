package main

import (
	"context"
	goflag "flag"
	"os"
	"os/signal"
	"syscall"

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
	for {
		select {
		case <-ctx.Done():
			done <- nil
			return
		}
	}
}

func root(cmd *cobra.Command, args []string) error {
	quit, cancel := context.WithCancel(context.TODO())
	done := make(chan error)
	go run(quit, done)

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
