package client

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

// RetryOnError will run the given resolver function and retry on rate limit exceeded errors
// or other retryable errors (like internal server errors) after waiting some amount of time.
func (c *Client) RetryOnError(ctx context.Context, tableName string, f func() error) error {
	retries := 0
	var err error
	for err = f(); retries < c.maxRetries; err = f() {
		if shouldRetry(err) {
			retryAfter := time.Duration((0.9 + rand.Float64()*0.2) * float64(c.backoff))
			retries++
			c.logger.Info().Str("table", tableName).Msgf("Got retryable error (%v), retrying in %.2fs (%d/%d)", err.Error(), retryAfter.Seconds(), retries, c.maxRetries)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(retryAfter):
				continue
			}
		}
		return err
	}
	return err
}

func shouldRetry(err error) bool {
	if err == nil {
		return false
	}
	if resp, ok := err.(*snyk.ErrorResponse); ok {
		return resp.Response.StatusCode >= http.StatusInternalServerError || resp.Response.StatusCode == http.StatusTooManyRequests
	}
	return false
}
