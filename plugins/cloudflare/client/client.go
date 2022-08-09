package client

import (
	"context"
	"errors"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
)

type AccountZones map[string]struct {
	AccountId string
	Zones     []string
}

type Clients map[string]Api

type Client struct {
	logger hclog.Logger

	accountsZones AccountZones
	clients       Clients

	ClientApi Api
	AccountId string
	ZoneId    string
}

const MaxItemsPerPage = 200

func New(logger hclog.Logger, clients Clients, clientApi Api, accountsZones AccountZones) Client {
	return Client{
		logger:        logger,
		accountsZones: accountsZones,
		clients:       clients,
		ClientApi:     clientApi,
	}
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

func (c *Client) withAccountId(accountId string) *Client {
	return &Client{
		logger:        c.logger.With("account_id", obfuscateId(accountId)),
		accountsZones: c.accountsZones,
		clients:       c.clients,
		ClientApi:     c.clients[accountId],
		AccountId:     accountId,
	}
}

func (c *Client) withZoneId(accountId, zoneId string) *Client {
	return &Client{
		logger:        c.logger.With("account_id", obfuscateId(accountId), "zone_id", obfuscateId(zoneId)),
		accountsZones: c.accountsZones,
		clients:       c.clients,
		ClientApi:     c.clients[accountId],
		AccountId:     accountId,
		ZoneId:        zoneId,
	}
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, diag.Diagnostics) {
	var diags diag.Diagnostics

	ctx := context.Background()
	providerConfig := config.(*Config)

	clientApi, err := getCloudflareClient(providerConfig)

	if err != nil {
		return nil, diags.Add(classifyError(err, diag.INTERNAL, diag.WithSeverity(diag.ERROR))) // TODO remove diag
	}

	var accountsZones = make(AccountZones)

	// Get available accounts
	accounts, _, err := clientApi.Accounts(ctx, cloudflare.AccountsListParams{})
	if err != nil {
		return nil, diags.Add(classifyError(err, diag.INTERNAL, diag.WithSeverity(diag.ERROR))) // TODO remove diag
	}

	for _, account := range accounts {
		// Get available zones  for each account
		zones, err := clientApi.ListZonesContext(ctx, cloudflare.WithZoneFilters("", account.ID, ""))
		if err != nil {
			// TODO log error and continue
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
		return nil, diags.Add(classifyError(errors.New("no accounts found"), diag.INTERNAL, diag.WithSeverity(diag.ERROR))) // TODO remove diag
	}

	clients := make(Clients)
	for _, account := range accountsZones {
		c, err := getCloudflareClient(providerConfig)
		if err != nil {
			return nil, diags.Add(classifyError(err, diag.INTERNAL, diag.WithSeverity(diag.ERROR))) // TODO remove diag
		}
		c.AccountID = account.AccountId
		clients[account.AccountId] = c
	}

	c := New(logger, clients, clientApi, accountsZones)
	return &c, nil
}

func getCloudflareClient(config *Config) (*cloudflare.API, error) {
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

func obfuscateId(accountId string) string {
	return accountId[:4] + "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
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
