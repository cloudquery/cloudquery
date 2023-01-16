package client

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/okta/okta-sdk-golang/v3/okta"
	"github.com/rs/zerolog"
	"github.com/thoas/go-funk"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger zerolog.Logger
	spec   specs.Source

	*okta.APIClient
}

const exampleDomain = "https://<CHANGE_THIS_TO_YOUR_OKTA_DOMAIN>.okta.com"

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.spec.Name
}

func New(logger zerolog.Logger, s specs.Source, okt *okta.APIClient) *Client {
	return &Client{
		APIClient: okt,

		logger: logger,
		spec:   s,
	}
}

func Configure(_ context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	oktaSpec := &Spec{}
	if err := s.UnmarshalSpec(oktaSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal okta spec: %w", err)
	}

	oktaToken, ok := os.LookupEnv("OKTA_API_TOKEN")
	if !ok {
		if oktaSpec.Token == "" {
			return nil, errors.New("missing OKTA_API_TOKEN, either set it as an environment variable or pass it in the configuration")
		}

		oktaToken = oktaSpec.Token
	}

	if oktaSpec.Domain == "" || oktaSpec.Domain == exampleDomain {
		return nil, errors.New(`failed to configure provider, please set your okta "domain" in okta.yml`)
	}

	cf := okta.NewConfiguration(
		okta.WithOrgUrl(oktaSpec.Domain),
		okta.WithToken(oktaToken),
		okta.WithCache(true),
	)
	c := okta.NewAPIClient(cf)

	return New(logger, s, c), nil
}

func ResolveNullableTime(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		data := funk.Get(resource.Item, path)
		if data == nil {
			return nil
		}
		ts, ok := data.(okta.NullableTime)
		if !ok {
			return fmt.Errorf("unexpected type, want \"okta.NullableTime\", have \"%T\"", data)
		}
		if !ts.IsSet() {
			return resource.Set(c.Name, nil)
		}
		return resource.Set(c.Name, ts.Get())
	}
}
