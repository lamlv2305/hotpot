package pond

import (
	"errors"
	"time"
)

var ErrNoResourcesAvailable = errors.New("no resources available")

type Option[T any] func(*pond[T])

func WithTimout[T any](timeout time.Duration) Option[T] {
	return func(p *pond[T]) {
		p.timeout = timeout
	}
}

func WithLocker[T any](locker Locker[T]) Option[T] {
	return func(p *pond[T]) {
		p.locker = locker
	}
}
