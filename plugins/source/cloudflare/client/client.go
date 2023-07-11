package client

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/v4/schema"
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
	idStrings := []string{
		c.AccountId,
		c.ZoneId,
	}

	return strings.TrimRight(strings.Join(idStrings, ":"), ":")
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

func Configure(ctx context.Context, logger zerolog.Logger, spec *Spec) (schema.ClientMeta, error) {
	clientApi, err := getCloudflareClient(spec)
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
		if len(spec.Accounts) > 0 && !funk.ContainsString(spec.Accounts, account.ID) {
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
		c, err := getCloudflareClient(spec)
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
