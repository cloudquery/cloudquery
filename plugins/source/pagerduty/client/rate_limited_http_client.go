package client

import (
	"net/http"

	"github.com/PagerDuty/go-pagerduty"
	"golang.org/x/time/rate"
)

// Note that this is an implementatioin of 'pagerduty.HTTPClient', not of 'http.Client'.
func newRateLimitedHttpClient(underlyingHttpClient pagerduty.HTTPClient, maxRequestsPerSecond int) *RateLimitedHttpClient {
	return &RateLimitedHttpClient{
		underlyingHttpClient: underlyingHttpClient,
		limiter: rate.NewLimiter(
			/*r=*/ rate.Limit(maxRequestsPerSecond),
			/*b=*/ 1,
		),
	}
}

// Note that this is an implementatioin of 'pagerduty.HTTPClient', not of 'http.Client'.
type RateLimitedHttpClient struct {
	underlyingHttpClient pagerduty.HTTPClient

	limiter *rate.Limiter
}

func (c *RateLimitedHttpClient) Do(req *http.Request) (*http.Response, error) {
	if err := c.limiter.Wait(req.Context()); err != nil {
		return nil, err
	}

	return c.underlyingHttpClient.Do(req)
}
