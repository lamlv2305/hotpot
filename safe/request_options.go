package safe

import "time"

type RequestOptions struct {
	delay   time.Duration
	timeout time.Duration
}

func WithDelay(delay time.Duration) func(*RequestOptions) {
	return func(r *RequestOptions) {
		r.delay = delay
	}
}

func WithTimeout(timeout time.Duration) func(*RequestOptions) {
	return func(r *RequestOptions) {
		r.timeout = timeout
	}
}
