package client

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client/services"
	"github.com/cloudquery/plugin-sdk/v4/schema"
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

func (*Client) ID() string {
	return "fastly"
}

func (c *Client) withServiceAndRegion(service *fastly.Service, region string) schema.ClientMeta {
	return &Client{
		logger:     c.logger.With().Str("service_id", service.ID).Str("Region", region).Logger(),
		Fastly:     c.Fastly,
		maxRetries: c.maxRetries,
		backoff:    c.backoff,
		Service:    service,
		Region:     region,
		services:   c.services,
		regions:    c.regions,
	}
}

func Configure(_ context.Context, logger zerolog.Logger, spec Spec) (schema.ClientMeta, error) {
	client, err := fastly.NewClient(spec.FastlyAPIKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create fastly client: %w", err)
	}

	fastlyServices, err := listServices(client, spec)
	if err != nil {
		return nil, fmt.Errorf("failed to list services: %w", err)
	}
	regions, err := listRegions(client)
	if err != nil {
		return nil, fmt.Errorf("failed to list regions: %w", err)
	}

	return &Client{
		logger:     logger,
		Fastly:     client,
		maxRetries: defaultMaxRetries,
		backoff:    defaultBackoff,
		services:   fastlyServices,
		regions:    regions,
		Spec:       spec,
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
