package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AzureDevOpsClient struct {
}

func (*AzureDevOpsClient) Logger() *zerolog.Logger {
	return &log.Logger
}

func (*AzureDevOpsClient) ID() string {
	return "AzureDevOpsClient"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	return &AzureDevOpsClient{}, nil
}
