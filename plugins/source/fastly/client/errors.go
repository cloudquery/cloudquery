package client

import (
	"context"
	"math/rand"
	"time"

	"github.com/fastly/go-fastly/v7/fastly"
)

// RetryOnError will run the given resolver function and retry on rate limit exceeded errors
// or other retryable errors (like internal server errors) after waiting some amount of time.
func (c *Client) RetryOnError(ctx context.Context, tableName string, f func() error) error {
	retries := 0
	for err := f(); err != nil; err = f() {
		if ferr, ok := err.(*fastly.HTTPError); ok && isRetryable(ferr) {
			retryAfter := time.Duration(rand.Float64() * float64(c.backoff))
			retries++
			c.logger.Info().Str("table", tableName).Msgf("Got retryable error (%v), retrying in %.2fs (%d/%d)", err, retryAfter.Seconds(), retries, c.maxRetries)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(retryAfter):
				continue
			}
		}
		return err
	}
	return nil
}

func isRetryable(err *fastly.HTTPError) bool {
	return err.StatusCode >= 500 || err.StatusCode == 429
}
