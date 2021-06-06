package client

import (
	"github.com/cloudquery/cq-provider-sdk/plugin/schema"
	"github.com/hashicorp/go-hclog"
)


type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger     hclog.Logger

	// Usually you store here your 3rd party clients and use them in the fetcher
	ThirdPartyClient interface{}
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}


func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	providerConfig := config.(*Config)
	_ = providerConfig
	// Init your client and 3rd party clients using the user's configuration
	// passed by the SDK providerConfig
	client := Client{
		logger:   logger,
		ThirdPartyClient: nil,
	}

	// Return the initialized client and it will be passed to your resources
	return &client, nil
}

