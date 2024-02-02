package client

import (
	_ "embed"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type (
	Spec struct {
		// Token for Okta API access.
		// You can set this with an `OKTA_API_TOKEN` environment variable.
		Token string `json:"token" jsonschema:"required,minLength=1"`

		// Specify the Okta domain you are fetching from.
		// [Visit this link](https://developer.okta.com/docs/guides/find-your-domain/findorg/) to find your Okta domain.
		Domain    string     `json:"domain" jsonschema:"required,pattern=^https?://[^\n<>]+\\.okta\\.com$"`
		RateLimit *RateLimit `json:"rate_limit"`

		// Enables debug logs within the Okta SDK.
		Debug bool `json:"debug,omitempty" jsonschema:"default=false"`

		// Number of concurrent requests to be made to Okta API.
		Concurrency int `json:"concurrency" jsonschema:"minimum=1,default=10000"`
	}
	RateLimit struct {
		// Max backoff interval to be used.
		MaxBackoff time.Duration `json:"max_backoff,omitempty" jsonschema:"minimum=30,default=30"`

		// Max retries to be performed.
		MaxRetries int32 `json:"max_retries,omitempty" jsonschema:"minimum=2,default=2"`
	}
)

const (
	OktaAPIToken = "OKTA_API_TOKEN"
)

//go:embed schema.json
var JSONSchema string

func (s *Spec) SetDefaults(logger *zerolog.Logger) {
	const (
		minRetries = int32(2)
		minBackOff = 30 * time.Second
	)

	if s.RateLimit == nil {
		s.RateLimit = new(RateLimit)
	}

	if s.RateLimit.MaxRetries < minRetries {
		s.RateLimit.MaxRetries = minRetries
	}

	if s.RateLimit.MaxBackoff < minBackOff {
		s.RateLimit.MaxBackoff = minBackOff
	}

	if len(s.Token) == 0 {
		logger.Warn().Msgf("usage of %q environment variable value is deprecated and will be dropped in a future release", OktaAPIToken)
		s.Token = os.Getenv(OktaAPIToken)
	}

	if s.Concurrency < 1 {
		s.Concurrency = 10000
	}
}

func (s Spec) Validate() error {
	if len(s.Token) == 0 {
		return fmt.Errorf("missing API token (should be set in the configuration or as %q environment variable)", OktaAPIToken)
	}

	const exampleDomain = "https://<CHANGE_THIS_TO_YOUR_OKTA_DOMAIN>.okta.com"
	switch s.Domain {
	case "", exampleDomain:
		return fmt.Errorf("missing \"domain\" in plugin configuration")
	}

	return nil
}
