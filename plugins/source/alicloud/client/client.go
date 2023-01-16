package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	services  map[string]map[string]*Services // account id -> region id -> Services
	logger    zerolog.Logger
	Spec      Spec
	AccountID string
	Region    string
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return strings.Join([]string{c.AccountID, c.Region}, ":")
}

func (c *Client) Services() *Services {
	return c.services[c.AccountID][c.Region]
}
func (c *Client) withAccountIDAndRegion(accountID, regionID string) *Client {
	return &Client{
		logger:    c.logger.With().Str("account_id", accountID).Str("region_id", regionID).Logger(),
		services:  c.services,
		AccountID: accountID,
		Region:    regionID,
	}
}

func New(_ context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	var spec Spec
	err := s.UnmarshalSpec(&spec)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal alicloud spec: %w", err)
	}
	if err := spec.Validate(); err != nil {
		return nil, err
	}

	services := make(map[string]map[string]*Services)
	for _, account := range spec.Accounts {
		for _, region := range account.RegionIDs {
			if _, ok := services[account.Name]; !ok {
				services[account.Name] = make(map[string]*Services)
			}
			services[account.Name][region], err = initServices(account, region)
			if err != nil {
				return nil, err
			}
		}
	}
	return &Client{logger: logger, Spec: spec, services: services}, nil
}

// mostly used for updating services in testing
func (c *Client) updateServices(svcs Services) {
	for accountID := range c.services {
		for regionID := range c.services[accountID] {
			c.services[accountID][regionID] = &svcs
		}
	}
}
