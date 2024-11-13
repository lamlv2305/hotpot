package pond

import (
	"context"
	"errors"
)

func NewPool[T any](resources ...T) *pool[T] {
	p := &pool[T]{
		resources: make(chan T, len(resources)),
	}

	// Fill resource to channel
	for idx := range resources {
		p.resources <- resources[idx]
	}

	return p
}

type pool[T any] struct {
	resources chan T
}

func (p *pool[T]) Acquire() (*Resource[T], error) {
	select {
	case res := <-p.resources:
		result := Resource[T]{
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

func (p *pool[T]) Do(ctx context.Context, task Task[T]) error {
	res, err := p.Acquire()
	if err != nil {
		return err
	}
	defer res.Release()

	var timeoutCtx context.Context
	var cancel context.CancelFunc

	if deadline := task.Deadline(); deadline != nil {
		timeoutCtx, cancel = context.WithDeadline(ctx, *deadline)
	} else {
		timeoutCtx, cancel = context.WithCancel(ctx)
	}

	defer cancel()

	return task.Do(timeoutCtx, res.value)
}

type Resource[T any] struct {
	value   T
	release func()
}

func (p *Resource[T]) Release() {
	p.release()
}
