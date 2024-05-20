package client

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/hermanschaaf/cq-source-xkcd/internal/xkcd"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger  zerolog.Logger
	XKCD    *xkcd.Client
	Backend state.Client
}

func (c *Client) ID() string {
	return "xkcd"
}

func New(logger zerolog.Logger, spec Spec, services *xkcd.Client, bk state.Client) schema.ClientMeta {
	return &Client{
		Logger:  logger,
		XKCD:    services,
		Backend: bk,
	}
}
