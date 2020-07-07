package dlock

import "context"

type Locker interface {
	Lock(ctx context.Context) error
	Unlock() error
}
