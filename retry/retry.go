package retry

import "context"

type RetryFunc func(ctx context.Context) error

func Do(ctx context.Context, b backoff, f RetryFunc) error