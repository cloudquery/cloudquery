package client

import "fmt"

type Spec struct {
	AccessToken string        `json:"access_token"`
	Orgs        []string      `json:"orgs"`
	Repos       []string      `json:"repos"`
	AppAuth     []AppAuthSpec `json:"app_auth"`
}

type AppAuthSpec struct {
	Org            string `json:"org"`
	AppID          string `json:"app_id"`
	PrivateKeyPath string `json:"private_key_path"`
	InstallationID string `json:"installation_id"`
}

func (s *Spec) Validate() error {
	if s.AccessToken == "" && len(s.AppAuth) == 0 {
		return fmt.Errorf("missing personal access token or app auth in configuration")
	}
	for _, appAuth := range s.AppAuth {
		if appAuth.Org == "" {
			return fmt.Errorf("missing org in app auth configuration")
		}
		if appAuth.AppID != "" && appAuth.PrivateKeyPath == "" {
			return fmt.Errorf("missing private key path in configuration")
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

func validateRepo(repo string) error {
	if repo == "" {
		return fmt.Errorf("missing repository")
	}
	if len(splitRepo(repo)) != 2 {
		return fmt.Errorf("invalid repository: %s (should be in <org>/<repo> format)", repo)
	}
	return nil
}
