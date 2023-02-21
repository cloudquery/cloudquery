package client

import (
	"net/http"

	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

func RatelimitedHttpClient(logger zerolog.Logger, hc *http.Client, limiter *rate.Limiter) *http.Client {
	c := *hc
	c.Transport = newLimitedTransport(logger, hc.Transport, limiter)
	return &c
}

type limitedTransport struct {
	logger    zerolog.Logger
	wrappedRT http.RoundTripper
	limiter   *rate.Limiter
}

func newLimitedTransport(logger zerolog.Logger, t http.RoundTripper, limiter *rate.Limiter) http.RoundTripper {
	if t == nil {
		t = http.DefaultTransport
	}
	return &limitedTransport{
		logger:    logger,
		wrappedRT: t,
		limiter:   limiter,
	}
}

func (t *limitedTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if !t.limiter.Allow() {
		t.logger.Debug().Msg("waiting for rate limiter...")
		err := t.limiter.Wait(r.Context())
		if err != nil {
			return nil, err
		}
	}

	return t.wrappedRT.RoundTrip(r)
}
