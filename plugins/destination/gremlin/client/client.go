package client

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers/batchwriter"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedSource
	logger zerolog.Logger
	spec   *Spec

	mu     sync.Mutex // protects client during session creation
	client *gremlingo.DriverRemoteConnection
	writer *batchwriter.BatchWriter
}

var AnonT = gremlingo.T__

func New(ctx context.Context, logger zerolog.Logger, spec []byte, _ plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", "gremlin").Logger(),
	}
	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal gremlin spec: %w", err)
	}
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	c.spec.SetDefaults()

	u := c.spec.Endpoint + "/gremlin"

	au, err := c.getAuthInfo(ctx, u)
	if err != nil {
		return nil, err
	}

	c.client, err = gremlingo.NewDriverRemoteConnection(u,
		func(settings *gremlingo.DriverRemoteConnectionSettings) {
			settings.TraversalSource = "g"
			settings.AuthInfo = au
			settings.MaximumConcurrentConnections = c.spec.MaxConcurrentConnections

			if logger.GetLevel() <= zerolog.TraceLevel {
				settings.LogVerbosity = gremlingo.Debug
			}

			if c.spec.Insecure {
				settings.TlsConfig = &tls.Config{InsecureSkipVerify: true}
			}
			settings.Logger = &Logger{Base: logger.With().Str("from", "gremlingo").Logger()}
		})
	if err != nil {
		return nil, err
	}

	c.writer, err = batchwriter.New(c, batchwriter.WithBatchSize(c.spec.BatchSize), batchwriter.WithBatchSizeBytes(c.spec.BatchSizeBytes))
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Client) Close(ctx context.Context) error {
	defer c.client.Close()
	if err := c.writer.Close(ctx); err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}
	return nil
}

func (c *Client) getAuthInfo(ctx context.Context, baseURL string) (gremlingo.AuthInfoProvider, error) {
	switch c.spec.AuthMode {
	case authModeNone:
		return nil, nil
	case authModeBasic:
		return gremlingo.BasicAuthInfo(c.spec.Username, c.spec.Password), nil

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

		gen := func() gremlingo.AuthInfoProvider {
			if err := signer.SignHTTP(ctx, cr, req, emptyStringSHA256, "neptune-db", c.spec.AWSRegion, time.Now()); err != nil {
				panic(err) // not ideal, but it's always nil
			}
			c.logger.Trace().Any("iam_headers", req.Header).Str("aws_region", c.spec.AWSRegion).Msg("IAM headers")
			return gremlingo.HeaderAuthInfo(req.Header)
		}
		return gremlingo.NewDynamicAuth(gen), nil

	default:
		return nil, fmt.Errorf("unhandled auth mode %q", c.spec.AuthMode)
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
