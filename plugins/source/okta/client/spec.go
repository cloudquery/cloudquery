package client

import (
	_ "embed"
	"fmt"
	"os"
	"errors"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/invopop/jsonschema"
	"github.com/rs/zerolog"
)

type (
	Spec struct {
		// Token for Okta API access.
		// You can set this with an `OKTA_API_TOKEN` environment variable.
		Token string `json:"token"`

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
		//  Max backoff interval to be used.
		// If the value specified is less than the default one, the default one is used.
		MaxBackoff configtype.Duration `json:"max_backoff,omitempty"`

		// Max retries to be performed.
		MaxRetries int32 `json:"max_retries,omitempty" jsonschema:"minimum=2,default=2"`
	}
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

	if s.RateLimit.MaxBackoff.Duration() < minBackOff {
		s.RateLimit.MaxBackoff = configtype.NewDuration(minBackOff)
	}

	if s.Concurrency < 1 {
		s.Concurrency = 10000
	}
}

func (s Spec) Validate() error {
	if len(s.Token) == 0 {
		return errors.New("missing \"token\" in plugin configuration")
	}

	const exampleDomain = "https://<CHANGE_THIS_TO_YOUR_OKTA_DOMAIN>.okta.com"
	switch s.Domain {
	case "", exampleDomain:
		return errors.New("missing \"domain\" in plugin configuration")
	}

	return nil
}

func (RateLimit) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.Properties.Value("max_backoff").Default = "30s"
}
