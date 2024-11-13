package pond

import (
	"context"
	"time"
)

type Task[T any] interface {
	Deadline() *time.Time
	Do(ctx context.Context, resource T) error
}
