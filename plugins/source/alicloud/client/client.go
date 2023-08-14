package client

import (
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/schema"
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
	return strings.Join([]string{"alicloud", c.AccountID, c.Region}, ":")
}

func (c *Client) Services() *Services {
	return c.services[c.AccountID][c.Region]
}

func (c *Client) WithAccountIDAndRegion(accountID, region string) *Client {
	return &Client{
		services:  c.services,
		logger:    c.logger.With().Str("account_id", accountID).Str("region", region).Logger(),
		Spec:      c.Spec,
		AccountID: accountID,
		Region:    region,
	}
}

func New(logger zerolog.Logger, spec Spec) (schema.ClientMeta, error) {
	var err error
	services := make(map[string]map[string]*Services)
	for _, account := range spec.Accounts {
		for _, region := range account.Regions {
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

// used for updating services in testing
func (c *Client) updateServices(svcs Services) {
	for accountID := range c.services {
		for region := range c.services[accountID] {
			c.services[accountID][region] = &svcs
		}
	}
}
