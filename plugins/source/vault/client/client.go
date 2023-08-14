package client

import (
	"github.com/hashicorp/vault/api"

	"github.com/rs/zerolog"
)

type Client struct {
	logger zerolog.Logger
	Spec   Spec

	VaultServices *Services
}

func (*Client) ID() string {
	return "vault"
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func New(logger zerolog.Logger, spec *Spec) (Client, error) {
	vaultClient, err := api.NewClient(&api.Config{
		Address: spec.VaultAddress,
	})
	if err != nil {
		return Client{}, err
	}

	return Client{
		logger:        logger,
		Spec:          *spec,
		VaultServices: NewServices(vaultClient),
	}, nil
}
