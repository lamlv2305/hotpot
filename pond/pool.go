package pond

import (
	"context"
	"time"
)

type Pond[T any] interface {
	Submit(ctx context.Context, task Task[T]) (any, error)
}

func NewPond[T any](resources []T, options ...Option[T]) *pond[T] {
	p := &pond[T]{
		timeout: time.Second * 30,
		locker:  &channelLocker[T]{},
	}

	for idx := range options {
		options[idx](p)
	}

	p.locker.Allocate(resources)

	return p
}

var _ Pond[any] = (*pond[any])(nil)

type pond[T any] struct {
	timeout time.Duration
	locker  Locker[T]
}

func (p *pond[T]) Submit(ctx context.Context, task Task[T]) (any, error) {
	timeout := p.timeout
	if task.Timeout() > 0 {
		timeout = task.Timeout()
	}

	deadline := time.Now().Add(timeout)

	timeoutCtx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	res, err := p.locker.Acquire(timeoutCtx)
	if err != nil {
		return nil, err
	}
	defer p.locker.Unlock(ctx, *res)

	return task.Run(timeoutCtx, *res)
}
