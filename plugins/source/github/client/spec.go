package client

import (
	_ "embed"
	"fmt"

	"github.com/invopop/jsonschema"
)

// Spec is the (nested) spec used by GitHub Source Plugin
type Spec struct {
	// Personal Access Token, required if not using App Authentication.
	AccessToken string `json:"access_token" jsonschema:"minLength=1"`
	// List of organizations to sync from. You must specify either orgs or repos in the configuration.
	Orgs []string `json:"orgs" jsonschema:"minItems=1"`
	// List of repositories to sync from. The format is owner/repo (e.g. cloudquery/cloudquery).
	// You must specify either orgs or repos in the configuration.
	Repos              []string            `json:"repos" jsonschema:"minItems=1,minLength=1,pattern=^[a-zA-Z0-9-_]+/[a-zA-Z0-9-_]+$"`
	AppAuth            []AppAuthSpec       `json:"app_auth" jsonschema:"minItems=1"`
	EnterpriseSettings *EnterpriseSettings `json:"enterprise"`

	// The best effort maximum number of Go routines to use.
	// Lower this number to reduce memory usage.
	// The plugin defaults to using the rate limit reported by the GitHub REST API `/rate_limit` endpoint, or 5000 if the rate limit is not available.
	Concurrency int `json:"concurrency,omitempty" jsonschema:"default=5000"`
	// Controls the number of parallel requests to GitHub when discovering repositories, a negative value means unlimited.
	DiscoveryConcurrency int `json:"discovery_concurrency,omitempty" jsonschema:"default=1"`
	// Include archived repositories when discovering repositories.
	IncludeArchivedRepos bool `json:"include_archived_repos,omitempty"`
}

type EnterpriseSettings struct {
	// The base URL of the GitHub Enterprise instance.
	BaseURL string `json:"base_url" jsonschema:"required,minLength=1"`
	// The upload URL of the GitHub Enterprise instance.
	UploadURL string `json:"upload_url" jsonschema:"required,minLength=1"`
}

type AppAuthSpec struct {
	// The GitHub organization to sync from.
	Org string `json:"org" jsonschema:"required,minLength=1"`
	// The GitHub App ID.
	AppID string `json:"app_id" jsonschema:"required,minLength=1"`
	// The path to the private key file used to authenticate the GitHub App.
	PrivateKeyPath string `json:"private_key_path" jsonschema:"minLength=1"`
	// The private key used to authenticate the GitHub App.
	PrivateKey string `json:"private_key" jsonschema:"minLength=1"`
	// The GitHub App installation ID.
	InstallationID string `json:"installation_id" jsonschema:"required,minLength=1"`
}

func (s *Spec) SetDefaults() {
	if s.DiscoveryConcurrency == 0 {
		s.DiscoveryConcurrency = 1
	}
}

func (s *Spec) Validate() error {
	if s.AccessToken == "" && len(s.AppAuth) == 0 {
		return fmt.Errorf("missing personal access token or app auth in configuration")
	}
	if s.EnterpriseSettings != nil {
		if err := s.ValidateEnterpriseSettings(); err != nil {
			return err
		}
	}
	for _, appAuth := range s.AppAuth {
		if appAuth.Org == "" {
			return fmt.Errorf("missing org in app auth configuration")
		}
		if appAuth.AppID != "" && (appAuth.PrivateKeyPath == "" && appAuth.PrivateKey == "") {
			return fmt.Errorf("missing private key specification in configuration. Please specify it using either `private_key` or `private_key_path`")
		}
		if appAuth.AppID != "" && (appAuth.PrivateKeyPath != "" && appAuth.PrivateKey != "") {
			return fmt.Errorf("both private key and private key path specified in configuration. Please remove the configuration for either `private_key_path` or `private_key`")
		}
		if appAuth.AppID != "" && appAuth.InstallationID == "" {
			return fmt.Errorf("missing installation id in configuration")
		}
	}
	if len(s.Orgs) == 0 && len(s.Repos) == 0 {
		return fmt.Errorf("missing orgs or repos in configuration")
	}
	for _, repo := range s.Repos {
		if err := validateRepo(repo); err != nil {
			return err
		}
	}
	return nil
}

func (s *Spec) ValidateEnterpriseSettings() error {
	if s.EnterpriseSettings.BaseURL == "" {
		return fmt.Errorf("enterprise base url is empty")
	}

	if s.EnterpriseSettings.UploadURL == "" {
		return fmt.Errorf("enterprise upload url is empty")
	}

	return nil
}

func validateRepo(repo string) error {
	if repo == "" {
		return fmt.Errorf("missing repository")
	}
	if len(splitRepo(repo)) != 2 {
		return fmt.Errorf("invalid repository: %s (should be in <org>/<repo> format)", repo)
	}
	return nil
}

func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.AllOf = []*jsonschema.Schema{
		{
			OneOf: []*jsonschema.Schema{
				{Required: []string{"access_token"}},
				{Required: []string{"app_auth"}},
			},
		},
		{
			OneOf: []*jsonschema.Schema{
				{Required: []string{"orgs"}},
				{Required: []string{"repos"}},
			},
		},
	}
}

func (AppAuthSpec) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.OneOf = []*jsonschema.Schema{
		{Required: []string{"private_key_path"}},
		{Required: []string{"private_key"}},
	}
}

//go:embed schema.json
var JSONSchema string
