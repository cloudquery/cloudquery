package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/specs"
	"github.com/rs/zerolog"
	"github.com/shenzhencenter/google-ads-pb/clients"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

const (
	_apiVer = "v13"

	// APIRef should be upgraded when the underlying API changes.
	// This also means that the table structure may change.
	APIRef = `https://developers.google.com/google-ads/api/reference/rpc/` + _apiVer

	// AdsScope is OAuth 2.0 scope for Google Ads (there's no readonly part)
	AdsScope = "https://www.googleapis.com/auth/adwords"
)

type Client struct {
	GoogleAdsClient *clients.GoogleAdsClient

	// Managed by multiplex
	CustomerID string
	ManagerID  string

	// key: management account ID, val: customer IDs.
	// See https://developers.google.com/google-ads/api/docs/migration/login-customer-id
	customers map[string][]string

	logger         zerolog.Logger
	developerToken string
}

func (c *Client) withManagerID(id string) *Client {
	res := *c
	res.ManagerID = id
	res.logger = res.logger.With().Str("manager_id", id).Logger()
	return &res
}

func (c *Client) withCustomerID(id string) *Client {
	res := *c
	res.CustomerID = id
	res.logger = res.logger.With().Str("customer_id", id).Logger()
	return &res
}

func (c *Client) ID() string {
	return "googleads:manager:{" + c.ManagerID + "}:customer:{" + c.CustomerID + "}"
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

var _ schema.ClientMeta = (*Client)(nil)

func Configure(ctx context.Context, logger zerolog.Logger, srcSpec specs.Source, options source.Options) (schema.ClientMeta, error) {
	spec := new(Spec)
	if err := srcSpec.UnmarshalSpec(&spec); err != nil {
		return nil, err
	}

	if err := spec.validate(); err != nil {
		return nil, err
	}

	opts := []option.ClientOption{
		option.WithUserAgent("cloudquery:source-googleads/" + srcSpec.Version),
		option.WithScopes(AdsScope),
		option.WithRequestReason("cloudquery resource fetch"),
		// we disable telemetry to boost performance and be on the safe side with telemetry
		option.WithTelemetryDisabled(),
	}

	if spec.OAuth != nil {
		tokenSource, err := spec.OAuth.getTokenSource(ctx, google.Endpoint, AdsScope)
		if err != nil {
			return nil, err
		}
		opts = append(opts, option.WithTokenSource(tokenSource))
	}

	client, err := clients.NewGoogleAdsClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	customerClient, err := clients.NewCustomerClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	defer customerClient.Close()

	c := &Client{
		GoogleAdsClient: client,
		logger:          logger.With().Str("plugin", "googleads").Logger(),
		developerToken:  spec.DeveloperToken,
	}

	return c, c.initCustomers(ctx, customerClient, spec)
}
