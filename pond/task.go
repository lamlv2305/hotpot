package pond

import (
	"context"
	"time"
)

type Task[T any] interface {
	Timeout() time.Duration
	Run(ctx context.Context, resource T) error
}
