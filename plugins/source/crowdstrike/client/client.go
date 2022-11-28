package client

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/crowdstrike/gofalcon/falcon"
	"github.com/crowdstrike/gofalcon/falcon/client"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Client struct {
	logger      zerolog.Logger
	CrowdStrike *client.CrowdStrikeAPISpecification
	spec        specs.Source
}

func (*Client) Logger() *zerolog.Logger {
	return &log.Logger
}

func (*Client) ID() string {
	return "CrowdStrikeClient"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	crowdStrikeSpec := &Spec{}
	if err := s.UnmarshalSpec(&crowdStrikeSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal CrowdStrike spec: %w", err)
	}
	clientId, ok := os.LookupEnv("FALCON_CLIENT_ID")
	if !ok {
		if crowdStrikeSpec.ClientID == "" {
			return nil, errors.New("missing FALCON_CLIENT_ID, either set it as an environment variable or pass it in the configuration")
		}
		clientId = crowdStrikeSpec.ClientID
	}

	secret, ok := os.LookupEnv("FALCON_CLIENT_SECRET")
	if !ok {
		if crowdStrikeSpec.ClientID == "" {
			return nil, errors.New("missing FALCON_CLIENT_SECRET, either set it as an environment variable or pass it in the configuration")
		}
		secret = crowdStrikeSpec.ClientSecret
	}

	c, _ := falcon.NewClient(&falcon.ApiConfig{
		ClientId:     clientId,
		ClientSecret: secret,
		Context:      ctx,
	})
	return &Client{
		logger:      logger,
		CrowdStrike: c,
		spec:        s,
	}, nil
}
