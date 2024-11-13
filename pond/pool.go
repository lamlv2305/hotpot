package pond

import (
	"context"
	"errors"
	"time"
)

type Pond[T any] interface {
	Do(ctx context.Context, task Task[T]) error
}

func NewPond[T any](resources []T, options ...Option[T]) *pond[T] {
	p := &pond[T]{
		resources: make(chan T, len(resources)),
		timeout:   time.Second * 30,
	}

	for idx := range options {
		options[idx](p)
	}

	// Fill resource to channel
	for idx := range resources {
		p.resources <- resources[idx]
	}

	return p
}

type pond[T any] struct {
	resources chan T
	timeout   time.Duration
}

func (p *pond[T]) Submit(ctx context.Context, task Task[T]) (err error) {
	timeout := p.timeout
	if task.Timeout() > 0 {
		timeout = task.Timeout()
	}

	deadline := time.Now().Add(timeout)

	timeoutCtx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	res, err := p.acquire(ctx)
	if err != nil {
		return err
	}
	defer res.release()

	return task.Run(timeoutCtx, res.value)
}

func (p *pond[T]) acquire(ctx context.Context) (*resource[T], error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	case res := <-p.resources:
		result := resource[T]{
			value: res,
			release: func() {
				select {
				case p.resources <- res:
				default:
					// Do nothing if the pool is already full
				}
			},
		}

		return &result, nil

	default:
		return nil, errors.New("no resources available")
	}
}

type resource[T any] struct {
	value   T
	release func()
}
