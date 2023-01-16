package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/backend"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/client"
	"golang.org/x/time/rate"
)

type Client struct {
	logger     zerolog.Logger
	sourceSpec specs.Source
	stSpec     Spec

	Services *client.API
	Backend  backend.Backend
}

func New(logger zerolog.Logger, sourceSpec specs.Source, stSpec Spec, services *client.API, bk backend.Backend) *Client {
	return &Client{
		logger:     logger,
		sourceSpec: sourceSpec,
		stSpec:     stSpec,
		Services:   services,
		Backend:    bk,
	}
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.sourceSpec.Name
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	stSpec := &Spec{}
	if err := s.UnmarshalSpec(stSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal stripe spec: %w", err)
	}

	services, err := getServiceClient(logger, stSpec)
	if err != nil {
		return nil, err
	}

	if err := validateAccess(services); err != nil {
		return nil, err
	}

	cl := New(logger, s, *stSpec, services, opts.Backend)
	return cl, nil
}

func getServiceClient(logger zerolog.Logger, spec *Spec) (*client.API, error) {
	if spec.APIKey == "" {
		return nil, errors.New("no api_key provided")
	}

	if !spec.StripeDebug {
		logger = logger.Level(zerolog.WarnLevel)
	}

	sCfg := &stripe.BackendConfig{
		LeveledLogger: &LeveledLogger{
			Logger: logger.With().Str("source", "stripe-client").Logger(),
		},
	}
	if spec.MaxRetries > 0 {
		sCfg.MaxNetworkRetries = stripe.Int64(spec.MaxRetries)
	}

	if spec.RateLimit < 1 {
		// https://stripe.com/docs/rate-limits
		if strings.HasPrefix(spec.APIKey, "sk_live") {
			spec.RateLimit = 90
		} else {
			spec.RateLimit = 20
		}
	}

	sCfg.HTTPClient = RatelimitedHttpClient(logger,
		&http.Client{Timeout: 80 * time.Second}, // default from stripe-go
		rate.NewLimiter(rate.Limit(spec.RateLimit), int(spec.RateLimit)),
	)

	c := &client.API{}
	c.Init(spec.APIKey, &stripe.Backends{
		API:     stripe.GetBackendWithConfig(stripe.APIBackend, sCfg),
		Connect: stripe.GetBackendWithConfig(stripe.ConnectBackend, sCfg),
		Uploads: stripe.GetBackendWithConfig(stripe.UploadsBackend, sCfg),
	})
	return c, nil
}

func validateAccess(svc *client.API) error {
	_, err := svc.Accounts.Get()
	if err != nil {
		return err
	}
	return nil
}
