package client

import (
	"errors"
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	Backends map[string]*TerraformBackend
	logger   hclog.Logger

	// CurrentBackend set by client multiplexer
	CurrentBackend string
}

func NewTerraformClient(logger hclog.Logger, backends map[string]*TerraformBackend) Client {
	return Client{
		Backends: backends,
		logger:   logger,
	}
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

func Configure(logger hclog.Logger, providerConfig interface{}) (schema.ClientMeta, diag.Diagnostics) {
	terraformConfig := providerConfig.(*Config)

	if len(terraformConfig.Config) == 0 {
		return nil, diag.FromError(errors.New("no config were provided"), diag.USER)
	}

	var backends = make(map[string]*TerraformBackend)
	for _, config := range terraformConfig.Config {
		config := config

		logger.Info("creating new backend", "type", config.BackendType)
		// create backend for each backend config
		if b, err := NewBackend(&config); err == nil { //nolint:revive
			backends[b.BackendName] = b
		} else {
			return nil, diag.FromError(fmt.Errorf("cannot initialize %s backend: %w", config.BackendType, err), diag.INTERNAL)
		}
	}

	client := NewTerraformClient(logger, backends)

	// Returns the initialized client with requested backends
	return &client, nil
}

func (c *Client) Backend() *TerraformBackend {
	if c.CurrentBackend != "" {
		backend := c.Backends[c.CurrentBackend]
		return backend
	}
	for _, backend := range c.Backends {
		return backend
	}
	return nil
}

// Sets the current backend to working with
func (c *Client) withSpecificBackend(backendName string) *Client {
	return &Client{
		Backends:       c.Backends,
		logger:         c.logger,
		CurrentBackend: backendName,
	}
}
