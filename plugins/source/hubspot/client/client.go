package client

import (
	"context"
	"github.com/clarkmcc/go-hubspot"
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

type Client struct {
	Authorizer  *hubspot.TokenAuthorizer
	Spec        spec.Spec
	RateLimiter *rate.Limiter
	// ObjectType is used for multiplexing when fetching `crm_pipelines`.
	ObjectType string
	Logger     zerolog.Logger
	Backend    state.Client
}

func (c *Client) ID() string {
	if c.ObjectType != "" {
		return "hubspot:" + c.ObjectType
	}
	return "hubspot"
}

func (c *Client) IsIncrementalSync() bool {
	return c.Backend != nil
}

// Used for multiplexing when fetching `crm_pipelines`
func (c *Client) withObjectType(objectType string) *Client {
	newClient := *c
	newClient.Logger = c.Logger.With().Str("object_type", objectType).Logger()
	newClient.ObjectType = objectType
	return &newClient
}

func New(_ context.Context, logger zerolog.Logger, s spec.Spec, backend state.Client) (schema.ClientMeta, error) {
	return &Client{
		Logger:     logger,
		Authorizer: hubspot.NewTokenAuthorizer(s.AppToken),
		Spec:       s,
		Backend:    backend,
		RateLimiter: rate.NewLimiter(
			/* r= */ rate.Limit(s.MaxRequestsPerSecond),
			/* b= */ 1,
		),
	}, nil
}
