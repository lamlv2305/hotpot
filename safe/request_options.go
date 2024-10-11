package safe

import "time"

type RequestOptions struct {
	waitDelay   time.Duration
	waitTimeout time.Duration
}

func WithWaitDelay(delay time.Duration) func(*RequestOptions) {
	return func(r *RequestOptions) {
		r.waitDelay = delay
	}
}

func WithWaitTimeout(timeout time.Duration) func(*RequestOptions) {
	return func(r *RequestOptions) {
		r.waitTimeout = timeout
	}
}
