package client

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"runtime"
	"time"

	"github.com/rs/zerolog"

	"github.com/cloudquery/pendo/pendo"
)

type Client struct {
	logger zerolog.Logger
	Spec   Spec
	Pendo  pendo.Client
}

func (*Client) ID() string {
	return "pendo"
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(_ context.Context, logger zerolog.Logger, s *Spec) (Client, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			ForceAttemptHTTP2:     true,
			MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		},
	}

	pendoClient, err := pendo.NewClient(
		pendo.WithAPIKey(s.PendoAPIKey),
		pendo.WithHttpDoer(httpClient),
		pendo.WithLogger(&logger),
	)
	if err != nil {
		return Client{}, fmt.Errorf("failed to create pendo client: %w", err)
	}

	c := Client{
		logger: logger,
		Spec:   *s,
		Pendo:  pendoClient,
	}

	return c, nil
}
