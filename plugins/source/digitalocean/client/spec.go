package client

import (
	_ "embed"
	"os"

	"github.com/pkg/errors"
)

// Spec defines DigitalOcean source plugin Spec
type Spec struct {
	// Token is the DigitalOcean API access token.
	//
	// If left empty it must be provided via either `DIGITALOCEAN_TOKEN` or `DIGITALOCEAN_ACCESS_TOKEN` environment variable.
	Token string `json:"token,omitempty" jsonschema_extras:"x-cq-auth=true"`
	// SpacesRegions is a list of DO regions to fetch spaces from, if not given we execute on all regions
	SpacesRegions []string `json:"spaces_regions,omitempty" jsonschema:"minLength=1"`
	// SpacesAccessKey is the secret access token generated in DO control panel
	SpacesAccessKey string `json:"spaces_access_key,omitempty" jsonschema_extras:"x-cq-auth=true"`
	// SpacesAccessKeyId is the unique identifier of the access key generated in the DO control panel
	SpacesAccessKeyId string `json:"spaces_access_key_id,omitempty" jsonschema_extras:"x-cq-auth=true"`
	// SpacesDebugLogging allows enabling AWS S3 request logging on spaces requests
	SpacesDebugLogging bool `json:"spaces_debug_logging,omitempty"`

	Concurrency int `json:"concurrency,omitempty" jsonschema:"minimum=1,default=10000"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency < 1 {
		s.Concurrency = 10000
	}

	if s.Token == "" {
		s.Token = getTokenFromEnv()
	}

	if s.SpacesAccessKey == "" || s.SpacesAccessKeyId == "" {
		s.SpacesAccessKeyId, s.SpacesAccessKey = getSpacesTokenFromEnv()
	}
}

func (s Spec) Validate() error {
	if s.Token == "" {
		return errors.New("missing API token")
	}
	return nil
}

func getTokenFromEnv() string {
	doToken := os.Getenv("DIGITALOCEAN_TOKEN")
	doAccessToken := os.Getenv("DIGITALOCEAN_ACCESS_TOKEN")
	if doToken != "" {
		return doToken
	}
	if doAccessToken != "" {
		return doAccessToken
	}
	return ""
}

func getSpacesTokenFromEnv() (string, string) {
	spacesAccessKey := os.Getenv("SPACES_ACCESS_KEY_ID")
	spacesSecretKey := os.Getenv("SPACES_SECRET_ACCESS_KEY")
	if spacesAccessKey == "" {
		return "", ""
	}
	if spacesSecretKey == "" {
		return "", ""
	}
	return spacesAccessKey, spacesSecretKey
}

//go:embed schema.json
var JSONSchema string
