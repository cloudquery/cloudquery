package client

import (
	"context"
	"errors"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
)

const MaxItemsPerPage = 200

type AccountZones map[string]struct {
	AccountId string
	Zones     []string
}

type Client struct {
	AccountsZones AccountZones
	AccountId     string
	ZoneId        string
	logger        hclog.Logger
	ClientApi     Api
}

func NewClient(log hclog.Logger, ClientApi Api, accountsZones AccountZones) Client {
	return Client{
		logger:        log,
		AccountsZones: accountsZones,
		ClientApi:     ClientApi,
	}
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

func getCloudflareClient(config *Config) (*cloudflare.API, error) {
	// Try to get the API token from the environment
	token := config.Token
	if token == "" {
		token = getApiTokenFromEnv()
	}

	if token != "" {
		clientApi, err := cloudflare.NewWithAPIToken(token)
		if err != nil {
			return nil, err
		}
		return clientApi, nil
	}

	apiKey := config.ApiKey
	apiEmail := config.ApiEmail

	if config.ApiKey == "" || config.ApiEmail == "" {
		apiKey, apiEmail = getApiKeyAndEmailFromEnv()
	}

	if apiKey != "" && apiEmail != "" {
		clientApi, err := cloudflare.New(apiKey, apiEmail)
		if err != nil {
			return nil, err
		}
		return clientApi, nil
	}

	return nil, errors.New("no API token or API key/email provided")
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

	c := NewClient(logger, clientApi, accountsZones)
	return &c, nil
}

func (c *Client) withAccountId(accountId string) *Client {
	return &Client{
		AccountsZones: c.AccountsZones,
		logger:        c.logger.With("account_id", obfuscateId(accountId)),
		AccountId:     accountId,
		ClientApi:     c.ClientApi,
	}
}

func (c *Client) withZoneId(AccountId, zoneId string) *Client {
	return &Client{
		AccountsZones: c.AccountsZones,
		logger:        c.logger.With("account_id", obfuscateId(AccountId), "zone_id", obfuscateId(zoneId)),
		AccountId:     AccountId,
		ZoneId:        zoneId,
		ClientApi:     c.ClientApi,
	}
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
