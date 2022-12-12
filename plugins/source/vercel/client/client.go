package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	logger     zerolog.Logger
	sharingID  string
	sourceSpec specs.Source
	veSpec     Spec

	Services *ServiceClient
}

func New(logger zerolog.Logger, sourceSpec specs.Source, veSpec Spec, services *ServiceClient) Client {
	return Client{
		logger:     logger,
		sourceSpec: sourceSpec,
		veSpec:     veSpec,
		Services:   services,
	}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.sourceSpec.Name
}

func Configure(_ context.Context, logger zerolog.Logger, s specs.Source) (schema.ClientMeta, error) {
	veSpec := &Spec{}
	if err := s.UnmarshalSpec(veSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal vercel spec: %w", err)
	}

	services, err := getServiceClient(veSpec)
	if err != nil {
		return nil, err
	}

	cl := New(logger, s, *veSpec, services)
	return &cl, nil
}

func getServiceClient(spec *Spec) (*ServiceClient, error) {
	if spec.AccessToken == "" {
		return nil, errors.New("no access token provided")
	}
	if spec.EndpointURL == "" {
		spec.EndpointURL = "https://api.vercel.com"
	}
	if spec.Timeout < 1 {
		spec.Timeout = 5
	}

	return &ServiceClient{
		Client: http.Client{
			Timeout: time.Duration(spec.Timeout) * time.Second,
		},

		baseURL: spec.EndpointURL,
		token:   spec.AccessToken,
		teamID:  spec.TeamID,
	}, nil
}
