package client

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/keplr-team/go-onfleet/onfleet"
	"github.com/rs/zerolog"
)

type Client struct {
	OnfleetClient *onfleet.Client

	Spec Spec

	// Goes into the `organization_id` column, to help distinguish orgs in case of multiple syncs from different orgs.
	OrganizationId string

	Logger zerolog.Logger
}

func (c *Client) ID() string {
	return "onfleet"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}

	authToken := os.Getenv("ONFLEET_API_KEY")
	if authToken == "" {
		return nil, fmt.Errorf("failed to get onfleet API key. please provide in via the 'ONFLEET_API_KEY' environment vairable (https://docs.onfleet.com/reference/authentication)")
	}

	onfleetClient := onfleet.Client{}
	onfleetClient.Init(authToken)

	orgId, err := getOrganizationId(ctx, &onfleetClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get org details: %s", err)
	}

	return &Client{
		OnfleetClient:  &onfleetClient,
		Spec:           pluginSpec,
		OrganizationId: orgId,
		Logger:         logger,
	}, nil
}

func getOrganizationId(ctx context.Context, onfleetClient *onfleet.Client) (string, error) {
	org, err := onfleetClient.Organization().Get(ctx)
	if err != nil {
		return "", err
	}

	if org.Id == "" {
		return "", fmt.Errorf("organization id is empty")
	}

	return org.Id, nil
}
