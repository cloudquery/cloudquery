package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client/services"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/fastly/go-fastly/v7/fastly"
	"github.com/rs/zerolog"
	"github.com/thoas/go-funk"
)

const (
	defaultMaxRetries = 5
	defaultBackoff    = 10 * time.Second
)

type Client struct {
	logger zerolog.Logger
	spec   specs.Source
	Fastly services.FastlyClient
	Spec   Spec

	// used for service-region multiplexing
	Service  *fastly.Service
	Region   string
	services []*fastly.Service
	regions  []string

	maxRetries int
	backoff    time.Duration // backoff duration between retries (jitter will be added)
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.spec.Name
}

func (c *Client) withServiceAndRegion(service *fastly.Service, region string) schema.ClientMeta {
	return &Client{
		logger:     c.logger.With().Str("service_id", service.ID).Str("Region", region).Logger(),
		spec:       c.spec,
		Fastly:     c.Fastly,
		maxRetries: c.maxRetries,
		backoff:    c.backoff,
		Service:    service,
		Region:     region,
		services:   c.services,
		regions:    c.regions,
	}
}

func Configure(ctx context.Context, logger zerolog.Logger, sourceSpec specs.Source, _ source.Options) (schema.ClientMeta, error) {
	var config Spec
	err := sourceSpec.UnmarshalSpec(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec: %w", err)
	}
	client, err := fastly.NewClient(config.FastlyAPIKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create fastly client: %w", err)
	}

	fastlyServices, err := listServices(client, config)
	if err != nil {
		return nil, fmt.Errorf("failed to list services: %w", err)
	}
	regions, err := listRegions(client)
	if err != nil {
		return nil, fmt.Errorf("failed to list regions: %w", err)
	}

	return &Client{
		logger:     logger,
		spec:       sourceSpec,
		Fastly:     client,
		maxRetries: defaultMaxRetries,
		backoff:    defaultBackoff,
		services:   fastlyServices,
		regions:    regions,
	}, nil
}

func listServices(client services.FastlyClient, cfg Spec) ([]*fastly.Service, error) {
	var fastlyServices []*fastly.Service
	p := client.NewListServicesPaginator(&fastly.ListServicesInput{
		PerPage: 100,
	})
	if p.HasNext() {
		s, err := p.GetNext()
		if err != nil {
			return nil, err
		}
		for _, service := range s {
			if len(cfg.Services) > 0 && !funk.ContainsString(cfg.Services, service.ID) {
				continue
			}
			fastlyServices = append(fastlyServices, service)
		}
	}
	return fastlyServices, nil
}

func listRegions(client services.FastlyClient) ([]string, error) {
	r, err := client.GetRegions()
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}
