package retry

import "time"

type Backoff interface {
	Next() (next time.Duration, stop bool)
}