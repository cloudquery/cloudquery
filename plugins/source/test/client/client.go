package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type TestClient struct {
}

func (*TestClient) Logger() *zerolog.Logger {
	return &log.Logger
}

func Configure(ctx context.Context, p *plugins.SourcePlugin, s specs.Source) (schema.ClientMeta, error) {
	return &TestClient{}, nil
}
