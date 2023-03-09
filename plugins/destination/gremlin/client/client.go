package client

import (
	"context"
	"crypto/tls"
	"fmt"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	destination.UnimplementedUnmanagedWriter
	destination.DefaultReverseTransformer
	logger     zerolog.Logger
	spec       specs.Destination
	pluginSpec Spec
	client     *gremlingo.DriverRemoteConnection
}

var AnonT = gremlingo.T__

func New(ctx context.Context, logger zerolog.Logger, destSpec specs.Destination) (destination.Client, error) {
	var err error
	c := &Client{
		logger: logger.With().Str("module", "gremlin").Logger(),
		spec:   destSpec,
	}
	var spec Spec
	if err := destSpec.UnmarshalSpec(&spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gremlin spec: %w", err)
	}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, err
	}

	c.pluginSpec = spec

	c.client, err = gremlingo.NewDriverRemoteConnection(spec.Endpoint+"/gremlin",
		func(settings *gremlingo.DriverRemoteConnectionSettings) {
			settings.TraversalSource = "g"
			settings.LogVerbosity = gremlingo.Debug

			if c.pluginSpec.Username != "" && c.pluginSpec.Password != "" {
				settings.AuthInfo = gremlingo.BasicAuthInfo(c.pluginSpec.Username, c.pluginSpec.Password)
			}

			if spec.Insecure {
				settings.TlsConfig = &tls.Config{InsecureSkipVerify: true}
			}
			settings.Logger = &Logger{Base: logger.With().Str("from", "gremlingo").Logger()}
		})
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Close(_ context.Context) error {
	c.client.Close()
	return nil
}
