package pond

import (
	"context"
)

type Locker[T any] interface {
	Allocate(resources []T)
	Acquire(ctx context.Context) (*T, error)
	Unlock(ctx context.Context, resource T) error
}

type Resource[T any] struct {
	Value   T
	Release func()
}

var _ Locker[any] = &channelLocker[any]{}

type channelLocker[T any] struct {
	resources chan T
}

func (c *channelLocker[T]) Allocate(resources []T) {
	c.resources = make(chan T, len(resources))

	for idx := range resources {
		c.resources <- resources[idx]
	}
}

// Acquire implements Locker.
func (c *channelLocker[T]) Acquire(ctx context.Context) (*T, error) {
	select {
	case res := <-c.resources:
		return &res, nil

	default:
		return nil, ErrNoResourcesAvailable
	}
}

// Unlock implements Locker.
func (c *channelLocker[T]) Unlock(ctx context.Context, resource T) error {
	select {
	case <-ctx.Done():
		return ctx.Err()

	case c.resources <- resource:
		return nil

	default:
		// Do nothing if the pool is already full
		return nil
	}
}
