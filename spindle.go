package dlock

import (
	"cloud.google.com/go/spanner"
	"github.com/flowerinthenight/spindle"
)

type SpindleLockOptions struct {
	Client   *spanner.Client
	Table    string // Spanner table name
	Name     string // lock name
	Id       string // optional, generated if empty
	Duration int64  // optional, will use spindle's default
}

func NewSpindleLock(opts *SpindleLockOptions) *SpindleLock {
	if opts == nil {
		return nil
	}

	s := &SpindleLock{opts: opts}
	sopts := []spindle.Option{}
	if opts.Id != "" {
		sopts = append(sopts, spindle.WithId(opts.Id))
	}

	if opts.Duration != 0 {
		sopts = append(sopts, spindle.WithDuration(opts.Duration))
	}

	s.lock = spindle.New(opts.Client, opts.Table, opts.Name, sopts...)
	return s
}

type SpindleLock struct {
	opts *SpindleLockOptions
	lock *spindle.Lock
}
