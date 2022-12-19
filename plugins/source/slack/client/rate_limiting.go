package client

import (
	"math/rand"
	"time"

	"github.com/slack-go/slack"
)

// RetryOnRateLimitError will run the given resolver function and retry on rate limit exceeded errors
// after waiting the specified amount of time.
func (c *Client) RetryOnRateLimitError(tableName string, f func() error) error {
	for err := f(); err != nil; err = f() {
		rateLimitedErr, ok := err.(*slack.RateLimitedError)
		if !ok {
			return err
		}
		randFloat64 := rand.Float64()
		jitter := time.Duration(randFloat64 * float64(1*time.Second))
		retryAfter := rateLimitedErr.RetryAfter + jitter
		c.logger.Info().Str("table", tableName).Msgf("Rate limit exceeded, retrying in %.2fs", retryAfter.Seconds())
		time.Sleep(retryAfter)
		continue
	}
	return nil
}
