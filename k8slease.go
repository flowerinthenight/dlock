package dlock

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
)

func k8sclient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

type K8sLockOption interface {
	Apply(*K8sLock)
}

type withK8sClient struct{ client *kubernetes.Clientset }

func (w withK8sClient) Apply(o *K8sLock) { o.client = w.client }

// WithK8sClient provides an option to set a k8s client object.
func WithK8sClient(c *kubernetes.Clientset) K8sLockOption { return withK8sClient{c} }

type withNamespace string

func (w withNamespace) Apply(o *K8sLock) { o.namespace = string(w) }

// WithNamespace provides an option to set the namespace value.
func WithNamespace(v string) K8sLockOption { return withNamespace(v) }

type withLeaseDuration time.Duration

func (w withLeaseDuration) Apply(o *K8sLock) { o.leaseDuration = time.Duration(w) }

// WithLeaseDuration provides an option to set the lease duration.
func WithLeaseDuration(v time.Duration) K8sLockOption { return withLeaseDuration(v) }

type withRenewDeadline time.Duration

func (w withRenewDeadline) Apply(o *K8sLock) { o.renewDeadline = time.Duration(w) }

// WithRenewDeadline provides an option to set the renew deadline.
func WithRenewDeadline(v time.Duration) K8sLockOption { return withRenewDeadline(v) }

type withRetryPeriod time.Duration

func (w withRetryPeriod) Apply(o *K8sLock) { o.retryPeriod = time.Duration(w) }

// WithRetryPeriod provides an option to set the renew deadline.
func WithRetryPeriod(v time.Duration) K8sLockOption { return withRetryPeriod(v) }

type withNewLeaderCallback func(string)

func (w withNewLeaderCallback) Apply(o *K8sLock) { o.onNewLeader = (func(string))(w) }

// WithNewLeaderCallback provides an option to set a callback when new lock is acquired.
func WithNewLeaderCallback(v func(string)) K8sLockOption { return withNewLeaderCallback(v) }

type withStartCallback func(context.Context)

func (w withStartCallback) Apply(o *K8sLock) { o.onStart = (func(context.Context))(w) }

// WithStartCallback provides an option to set a callback after we acquired the new lock.
func WithStartCallback(v func(context.Context)) K8sLockOption { return withStartCallback(v) }

// NewK8sLock returns an object that can be used to acquired/release a lock using k8s'
// LeaseLock resource object.
func NewK8sLock(id, name string, opts ...K8sLockOption) *K8sLock {
	lock := &K8sLock{
		id:            id,
		name:          name,
		namespace:     "default",
		leaseDuration: time.Minute,
		renewDeadline: 15 * time.Second,
		retryPeriod:   5 * time.Second,
	}

	// Apply provided options, if any.
	for _, opt := range opts {
		opt.Apply(lock)
	}

	return lock
}

type K8sLock struct {
	client        *kubernetes.Clientset
	id            string
	name          string
	namespace     string
	leaseDuration time.Duration
	renewDeadline time.Duration
	retryPeriod   time.Duration
	onNewLeader   func(string)
	onStart       func(context.Context)
	onStop        func()

	quit   context.Context
	cancel context.CancelFunc
}

// Lock attempts to acquire a k8s lock using LeaseLock. This call will unblock/return
// when lock is acquired or when ctx expires or is cancelled.
func (l *K8sLock) Lock(ctx context.Context) error {
	var err error
	if l.client == nil {
		l.client, err = k8sclient()
		if err != nil {
			return err
		}
	}

	// For the Unlock method.
	l.quit, l.cancel = context.WithCancel(context.TODO())
	leadch := make(chan struct{}, 1)

	go func() {
		lock := &resourcelock.LeaseLock{
			LeaseMeta: metav1.ObjectMeta{
				Name:      l.name,
				Namespace: l.namespace,
			},
			Client:     l.client.CoordinationV1(),
			LockConfig: resourcelock.ResourceLockConfig{Identity: l.id},
		}

		// Start the leader election code loop.
		leaderelection.RunOrDie(l.quit, leaderelection.LeaderElectionConfig{
			Lock:            lock,
			ReleaseOnCancel: true,
			LeaseDuration:   l.leaseDuration,
			RenewDeadline:   l.renewDeadline,
			RetryPeriod:     l.retryPeriod,
			Callbacks: leaderelection.LeaderCallbacks{
				OnStartedLeading: func(ctx context.Context) {
					leadch <- struct{}{}
					if l.onStart != nil {
						l.onStart(ctx)
					}
				},
				OnStoppedLeading: func() {
					if l.onStop != nil {
						l.onStop()
					}
				},
				OnNewLeader: func(identity string) {
					if l.onNewLeader != nil {
						l.onNewLeader(identity)
					}
				},
			},
		})
	}()

	select {
	case <-leadch:
	case <-ctx.Done():
	case <-l.quit.Done():
	}

	return nil
}

// Unlock releases the lock object.
func (l *K8sLock) Unlock() error {
	l.cancel()
	return nil
}
