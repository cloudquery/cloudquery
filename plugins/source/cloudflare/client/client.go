package client

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/thoas/go-funk"
)

type AccountZones map[string]struct {
	AccountId string
	Zones     []string
}

type Clients map[string]Api

type Client struct {
	logger zerolog.Logger

	accountsZones AccountZones
	clients       Clients

	ClientApi Api
	AccountId string
	ZoneId    string
}

const MaxItemsPerPage = 200

func New(logger zerolog.Logger, clients Clients, clientApi Api, accountsZones AccountZones) Client {
	return Client{
		logger:        logger,
		accountsZones: accountsZones,
		clients:       clients,
		ClientApi:     clientApi,
	}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.AccountId
}

func (c *Client) withAccountID(accountId string) *Client {
	return &Client{
		logger:        c.logger.With().Str("account_id", accountId).Logger(),
		accountsZones: c.accountsZones,
		clients:       c.clients,
		ClientApi:     c.clients[accountId],
		AccountId:     accountId,
	}
}

func (c *Client) withZoneID(accountId, zoneId string) *Client {
	return &Client{
		logger:        c.logger.With().Str("account_id", accountId).Str("zone_id", zoneId).Logger(),
		accountsZones: c.accountsZones,
		clients:       c.clients,
		ClientApi:     c.clients[accountId],
		AccountId:     accountId,
		ZoneId:        zoneId,
	}
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	cfSpec := &Spec{}
	if err := s.UnmarshalSpec(cfSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cloudflare spec: %w", err)
	}

	clientApi, err := getCloudflareClient(cfSpec)
	if err != nil {
		return nil, err
	}

	var accountsZones = make(AccountZones)

	// Get available accounts
	accounts, _, err := clientApi.Accounts(ctx, cloudflare.AccountsListParams{})
	if err != nil {
		return nil, err
	}

	for _, account := range accounts {
		if len(cfSpec.Accounts) > 0 && !funk.ContainsString(cfSpec.Accounts, account.ID) {
			continue
		}

		// Get available zones  for each account
		zones, err := clientApi.ListZonesContext(ctx, cloudflare.WithZoneFilters("", account.ID, ""))
		if err != nil {
			logger.Warn().Err(err).Msg("ListZonesContext failed")
			continue
		}
		var zoneIds []string
		for _, zone := range zones.Result {
			zoneIds = append(zoneIds, zone.ID)
		}

		accountsZones[account.ID] = struct {
			AccountId string
			Zones     []string
		}{
			AccountId: account.ID,
			Zones:     zoneIds,
		}
	}

	if len(accountsZones) == 0 {
		return nil, errors.New("no accounts found")
	}

	clients := make(Clients)
	for _, account := range accountsZones {
		c, err := getCloudflareClient(cfSpec)
		if err != nil {
			return nil, err
		}
		c.AccountID = account.AccountId
		clients[account.AccountId] = c
	}

	c := New(logger, clients, clientApi, accountsZones)
	return &c, nil
}

func getCloudflareClient(config *Spec) (*cloudflare.API, error) {
	// Try to get the API token from the environment
	token := config.Token
	if token == "" {
		token = getApiTokenFromEnv()
	}

	if token != "" {
		return cloudflare.NewWithAPIToken(token)
	}

	apiKey, apiEmail := config.ApiKey, config.ApiEmail

	if config.ApiKey == "" || config.ApiEmail == "" {
		apiKey, apiEmail = getApiKeyAndEmailFromEnv()
	}

	if apiKey != "" && apiEmail != "" {
		return cloudflare.New(apiKey, apiEmail)
	}

	return nil, errors.New("no API token or API key/email provided")
}

func getApiTokenFromEnv() string {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if apiToken == "" {
		return ""
	}
	return apiToken
}

func getApiKeyAndEmailFromEnv() (string, string) {
	apiKey := os.Getenv("CLOUDFLARE_API_KEY")
	apiEmail := os.Getenv("CLOUDFLARE_EMAIL")
	if apiKey == "" || apiEmail == "" {
		return "", ""
	}
	return apiKey, apiEmail
}
