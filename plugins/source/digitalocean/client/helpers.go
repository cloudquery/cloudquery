package client

import (
	"time"

	"github.com/digitalocean/godo"
)

// ThrottleFunc is passed to throttle wrapper
type ThrottleFunc func() error

// IsLimitReached - checks if API error is request limit error
func IsLimitReached(err error) bool {
	er, ok := err.(*godo.ErrorResponse)
	if !ok {
		return false
	}
	return er.Message == "Too many requests"
}

// ThrottleWrapper does API request until it is succeeded
func ThrottleWrapper(client *Client, doFunc ThrottleFunc) error {
	for {
		//todo maybe add a counter for max retries
		err := doFunc()
		if err != nil {
			if IsLimitReached(err) {
				client.Logger().Warn("API Rate limit exceeded. Request will be executed again after throttling delay")
				rate := client.DoClient.GetRate()
				client.Logger().Debug("Current API rate limits", "limit", rate.Limit, "remaining", rate.Remaining, "reset", rate.Reset.Time)
				// todo discover optimal delay
				time.Sleep(time.Second + 10)
				continue
			}
			return err
		}
		break
	}
	return nil
}
