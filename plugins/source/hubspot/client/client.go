package client

import (
	"context"
	"fmt"
	"os"

	"github.com/clarkmcc/go-hubspot"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

// Empirically tested that this is the largest page size that HubSpot allows.
const DefaultPageSize = 100

type Client struct {
	Authorizer *hubspot.TokenAuthorizer

	Spec Spec

	RateLimiter *rate.Limiter

	// Used for multiplexing when fetching `crm_pipelines`.
	ObjectType string

	Logger zerolog.Logger
}

func (c *Client) ID() string {
	if c.ObjectType != "" {
		return "hubspot:" + c.ObjectType
	}
	return "hubspot"
}

// Used for multiplexing when fetching `crm_pipelines`
func (c *Client) withObjectType(objectType string) *Client {
	newClient := *c
	newClient.Logger = c.Logger.With().Str("object_type", objectType).Logger()
	newClient.ObjectType = objectType
	return &newClient
}

func New(ctx context.Context, logger zerolog.Logger, s Spec) (schema.ClientMeta, error) {
	authToken := os.Getenv("HUBSPOT_APP_TOKEN")
	if authToken == "" {
		return nil, fmt.Errorf("failed to get hubspot auth token. Please provide an auth-token (see https://www.cloudquery.io/docs/plugins/sources/hubspot/overview#authentication)")
	}

	return &Client{
		Logger:     logger,
		Authorizer: hubspot.NewTokenAuthorizer(authToken),
		Spec:       s,
		RateLimiter: rate.NewLimiter(
			/* r= */ rate.Limit(*s.MaxRequestsPerSecond),
			/* b= */ 1,
		),
	}, nil
}
