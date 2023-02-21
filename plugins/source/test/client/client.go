package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/plugins/source"
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

func (*TestClient) ID() string {
	return "TestClient"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	return &TestClient{}, nil
}
