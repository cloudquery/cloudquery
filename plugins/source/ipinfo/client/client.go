package client

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/ipinfo/internal/ipinfo"
	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger
	Spec   Spec
	IPInfo *ipinfo.Client
}

func (c *Client) ID() string {
	// TODO: Change to either your plugin name or a unique dynamic identifier
	return "ipinfo"
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(ctx context.Context, logger zerolog.Logger, s *Spec) (Client, error) {
	// TODO: Add your client initialization here
	c, err := ipinfo.NewClient(s.Token)
	if err != nil {
		return Client{}, err
	}

	return Client{
		logger: logger,
		Spec:   *s,
		IPInfo: c,
	}, nil

}
