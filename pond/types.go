package pond

import "time"

type Option[T any] func(*pond[T])

func WithTimout[T any](timeout time.Duration) Option[T] {
	return func(p *pond[T]) {
		p.timeout = timeout
	}
}
