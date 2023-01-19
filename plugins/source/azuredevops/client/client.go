package client

import (
	"context"
	"fmt"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v6"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Client struct {
	logger zerolog.Logger
	*azuredevops.Connection
}

func (*Client) Logger() *zerolog.Logger {
	return &log.Logger
}

func (*Client) ID() string {
	return "AzureDevOpsClient"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	var spec Spec
	err := s.UnmarshalSpec(&spec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal GitHub spec: %w", err)
	}

	if spec.PersonalAccessToken == "" {
		return nil, fmt.Errorf("missing personal access token in configuration")
	}
	if spec.OrganizationURL == "" {
		return nil, fmt.Errorf("missing organization url in configuration")
	}

	connection := azuredevops.NewPatConnection(spec.OrganizationURL, spec.PersonalAccessToken)
	return &Client{
		logger:     logger,
		Connection: connection,
	}, nil
}
