package client

import (
	"context"
	"math/rand"
	"time"

	"github.com/slack-go/slack"
)

// RetryOnError will run the given resolver function and retry on rate limit exceeded errors
// or other retryable errors (like internal server errors) after waiting some amount of time.
func (c *Client) RetryOnError(ctx context.Context, tableName string, f func() error) error {
	retries := 0
	for err := f(); err != nil; err = f() {
		var retryAfter time.Duration
		randFloat64 := rand.Float64()
		switch typed := err.(type) {
		case *slack.RateLimitedError:
			jitter := time.Duration(randFloat64 * float64(c.backoff))
			retryAfter = typed.RetryAfter + jitter
			c.logger.Info().Str("table", tableName).Msgf("Rate limit exceeded, retrying in %.2fs", retryAfter.Seconds())
		default:
			if !isRetryable(err) || retries >= c.maxRetries {
				return err
			}
			retryAfter = time.Duration(randFloat64 * float64(c.backoff))
			retries++
			c.logger.Info().Str("table", tableName).Msgf("Got retryable error (%v), retrying in %.2fs (%d/%d)", err, retryAfter.Seconds(), retries, c.maxRetries)
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(retryAfter):
			continue
		}
	}
	return nil
}

func isRetryable(err error) bool {
	switch c := err.(type) {
	case *slack.StatusCodeError:
		return c.Retryable()
	case slack.StatusCodeError:
		return c.Retryable()
	}
	return false
}
