package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
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
	u := spec.Endpoint + "/gremlin"

	au, err := c.getAuthInfo(ctx, u)
	if err != nil {
		return nil, err
	}

	c.client, err = gremlingo.NewDriverRemoteConnection(u,
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

func (c *Client) getAuthInfo(ctx context.Context, baseURL string) (*gremlingo.AuthInfo, error) {
	switch c.pluginSpec.AuthMode {
	case authModeBasic:
		if c.pluginSpec.Username == "" && c.pluginSpec.Password == "" {
			return nil, nil
		}
		return gremlingo.BasicAuthInfo(c.pluginSpec.Username, c.pluginSpec.Password), nil

	case authModeIAM:
		req, err := http.NewRequest(http.MethodGet, baseURL, strings.NewReader(""))
		if err != nil {
			return nil, err
		}
		signer := v4.NewSigner(credentials.NewSharedCredentials("", ""))
		if _, err := signer.Sign(req, strings.NewReader(""), "neptune-db", "us-east-1", time.Now()); err != nil {
			return nil, err
		}

		c.logger.Trace().Any("iam_headers", req.Header).Msg("IAM headers")
		return gremlingo.HeaderAuthInfo(req.Header), nil

	default:
		return nil, fmt.Errorf("unhandled auth mode %q", c.pluginSpec.AuthMode)
	}
}
