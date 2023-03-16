package retry

import "context"

type RetryFunc func(ctx context.Context) error

type retryableError struct {
	err error
}



func Do(ctx context.Context, b Backoff, f RetryFunc) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		err := f(ctx)
		if err == nil {
			return nil
		}

	}
}
