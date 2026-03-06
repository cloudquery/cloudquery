package client

import (
	"github.com/cloudquery/cloudquery/plugins/source/tenable/internal/tenable"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger  zerolog.Logger
	Tenable *tenable.Client
	Backend state.Client
	Spec    Spec
}

func (*Client) ID() string {
	return "tenable"
}

func New(logger zerolog.Logger, spec Spec, tenableClient *tenable.Client, backend state.Client) schema.ClientMeta {
	return &Client{
		Logger:  logger,
		Tenable: tenableClient,
		Backend: backend,
		Spec:    spec,
	}
}
