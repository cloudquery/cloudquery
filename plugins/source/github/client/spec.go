package client

import "fmt"

// Spec is the (nested) spec used by GitHub Source Plugin
type Spec struct {
	// Personal Access Token, required if not using App Authentication.
	AccessToken string `json:"access_token"`
	// List of organizations to sync from. You must specify either orgs or repos in the configuration.
	Orgs []string `json:"orgs"`
	// List of repositories to sync from. The format is owner/repo (e.g. cloudquery/cloudquery).
	// You must specify either orgs or repos in the configuration.
	Repos              []string            `json:"repos"`
	AppAuth            []AppAuthSpec       `json:"app_auth"`
	EnterpriseSettings *EnterpriseSettings `json:"enterprise"`

	// The best effort maximum number of Go routines to use.
	// Lower this number to reduce memory usage.
	Concurrency int `json:"concurrency,omitempty" jsonschema:"default=10000"`
	// Controls the number of parallel requests to GitHub when discovering repositories, a negative value means unlimited.
	DiscoveryConcurrency int `json:"discovery_concurrency,omitempty" jsonschema:"default=1"`
	// Skip archived repositories when discovering repositories.
	SkipArchivedRepos bool `json:"skip_archived_repos,omitempty"`
}

type EnterpriseSettings struct {
	BaseURL   string `json:"base_url"`
	UploadURL string `json:"upload_url"`
}

type AppAuthSpec struct {
	Org            string `json:"org"`
	AppID          string `json:"app_id"`
	PrivateKeyPath string `json:"private_key_path"`
	PrivateKey     string `json:"private_key"`
	InstallationID string `json:"installation_id"`
}

func (s *Spec) SetDefaults() {
	if s.Concurrency == 0 {
		s.Concurrency = 10000
	}
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
