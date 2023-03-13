package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
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

	mu     sync.Mutex // protects client during session creation
	client *gremlingo.DriverRemoteConnection
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
			settings.AuthInfo = au
			settings.MaximumConcurrentConnections = spec.MaxConcurrentConnections

			if logger.GetLevel() <= zerolog.DebugLevel {
				settings.LogVerbosity = gremlingo.Debug
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

func (c *Client) getAuthInfo(ctx context.Context, baseURL string) (*gremlingo.AuthInfo, error) {
	switch c.pluginSpec.AuthMode {
	case authModeNone:
		return nil, nil
	case authModeBasic:
		return gremlingo.BasicAuthInfo(c.pluginSpec.Username, c.pluginSpec.Password), nil

	case authModeAWS:
		// emptyStringSHA256 is a SHA256 of an empty string
		const emptyStringSHA256 = `e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855`

		req, err := http.NewRequest(http.MethodGet, baseURL, strings.NewReader(""))
		if err != nil {
			return nil, err
		}

		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
		}
		cr, err := cfg.Credentials.Retrieve(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve AWS credentials: %w", err)
		}
		signer := v4.NewSigner()
		if err := signer.SignHTTP(ctx, cr, req, emptyStringSHA256, "neptune-db", c.pluginSpec.AWSRegion, time.Now()); err != nil {
			return nil, err
		}
		c.logger.Trace().Any("iam_headers", req.Header).Str("aws_region", c.pluginSpec.AWSRegion).Msg("IAM headers")
		return gremlingo.HeaderAuthInfo(req.Header), nil

	default:
		return nil, fmt.Errorf("unhandled auth mode %q", c.pluginSpec.AuthMode)
	}
}

func (c *Client) newSession() (*gremlingo.DriverRemoteConnection, func(), error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	sess, err := c.client.CreateSession()
	if err != nil {
		return nil, nil, err
	}

	return sess, func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		sess.Close()
	}, nil
}
