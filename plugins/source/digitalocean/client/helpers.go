package client

import (
	"context"
	"errors"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/digitalocean/godo"
)

const MAX_RETRIES = 5

// IsLimitReached - checks if API error is request limit error
func IsLimitReached(err error) bool {
	unwrapped := errors.Unwrap(err)
	er, ok := unwrapped.(*godo.ErrorResponse)
	if !ok {
		return false
	}
	return er.Message == "Too many requests"
}

// ThrottleWrapper does API request until it is succeeded
func ThrottleWrapper(ctx context.Context, client *Client, doFunc retry.RetryableFunc) error {
	err := retry.Do(
		doFunc,
		retry.OnRetry(func(n uint, err error) {
			client.Logger().Warn().Uint("retry", n).Msg("API Rate limit exceeded. Request will be executed again after throttling delay")
			rate := client.DoClient.GetRate()
			client.Logger().Debug().Int("limit", rate.Limit).Int("remaining", rate.Remaining).Time("reset", rate.Reset.Time).Msg("Current API rate limits")
		}),
		retry.RetryIf(IsLimitReached),
		retry.Attempts(MAX_RETRIES),
		retry.Context(ctx),
		retry.Delay(time.Second+5), // todo discover optimal delay
	)
	return err
}
