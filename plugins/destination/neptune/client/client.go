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
		logger: logger.With().Str("module", "neptune").Logger(),
		spec:   destSpec,
	}
	var spec Spec
	if err := destSpec.UnmarshalSpec(&spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal neptune spec: %w", err)
	}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		return nil, err
	}

	c.pluginSpec = spec

	au, err := c.getAuthInfo(ctx)
	if err != nil {
		return nil, err
	}

	c.client, err = gremlingo.NewDriverRemoteConnection(spec.Endpoint+"/gremlin",
		func(settings *gremlingo.DriverRemoteConnectionSettings) {
			settings.TraversalSource = "g"
			settings.LogVerbosity = gremlingo.Debug
			settings.AuthInfo = au

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

func (c *Client) getAuthInfo(ctx context.Context) (*gremlingo.AuthInfo, error) {
	if c.pluginSpec.Username != "" && c.pluginSpec.Password != "" {
		return gremlingo.BasicAuthInfo(c.pluginSpec.Username, c.pluginSpec.Password), nil
	}

	return nil, nil
	/*
		cfg, err := config.LoadDefaultConfig(ctx, config.WithDefaultRegion("us-east-1"))
		if err != nil {
			return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
		}
		cr, err := cfg.Credentials.Retrieve(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve AWS credentials: %w", err)
		}
		sig := signer.NewSigner()
		u := c.pluginSpec.Endpoint + "/gremlin"
		//u = strings.ReplaceAll(u, "wss://", "ws://")
		rq, _ := http.NewRequest(http.MethodGet, u, nil)
		if err := sig.SignHTTP(ctx, cr, rq, "", "neptune", cfg.Region, time.Now()); err != nil {
			return nil, fmt.Errorf("unsable to sign request: %w", err)
		}

		hdr := rq.Header
		c.logger.Debug().Any("auth_headers", hdr).Str("url", rq.URL.String()).Msg("signed headers")
		return gremlingo.HeaderAuthInfo(hdr), nil
	*/
}
