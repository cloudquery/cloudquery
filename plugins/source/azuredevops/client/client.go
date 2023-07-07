package client

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
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

func New(logger zerolog.Logger, spec Spec) (schema.ClientMeta, error) {
	connection := azuredevops.NewPatConnection(spec.OrganizationURL, spec.PersonalAccessToken)
	return &Client{
		logger:     logger,
		Connection: connection,
	}, nil
}
