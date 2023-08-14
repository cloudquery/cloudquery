package client

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/homebrew/internal/homebrew"
)

// RetryOnError will run the given resolver function and retry on rate limit exceeded errors
// or other retryable errors (like internal server errors) after waiting some amount of time.
func (c *Client) RetryOnError(ctx context.Context, tableName string, f func() error) error {
	retries := 0
	for err := f(); err != nil; err = f() {
		if isRetryable(err) {
			retryAfter := time.Duration(rand.Float64() * float64(c.Backoff))
			retries++
			c.Logger.Info().Str("table", tableName).Msgf("Got retryable error (%v), retrying in %.2fs (%d/%d)", err, retryAfter.Seconds(), retries, c.MaxRetries)
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

func isRetryable(err error) bool {
	var httpErr homebrew.HTTPError
	if errors.As(err, &httpErr) {
		return httpErr.Code >= http.StatusInternalServerError || httpErr.Code == http.StatusTooManyRequests
	}
	return false
}
