package client

import "fmt"

type Spec struct {
	ClientId       string   `json:"client_id"`
	ClientSecret   string   `json:"client_secret"`
	Username       string   `json:"username"`
	Password       string   `json:"password"`
	SFAPIVersion   string   `json:"api_version"`
	IncludeObjects []string `json:"include_objects"`
	ExcludeObjects []string `json:"exclude_objects"`
	OAuthURL       string   `json:"oauth_url"`
}

func (s *Spec) Validate() error {
	if s.ClientId == "" {
		return fmt.Errorf("client_id is required")
	}
	if s.ClientSecret == "" {
		return fmt.Errorf("client_secret is required")
	}
	if s.Username == "" {
		return fmt.Errorf("username is required")
	}
	if s.Password == "" {
		return fmt.Errorf("password is required")
	}
	if s.SFAPIVersion == "" {
		return fmt.Errorf("api_version is required")
	}
	return nil
}

func (s *Spec) SetDefaults() {
	if s.IncludeObjects == nil {
		s.IncludeObjects = []string{"*"}
	}
	if s.ExcludeObjects == nil {
		s.ExcludeObjects = []string{}
	}
	if s.OAuthURL == "" {
		s.OAuthURL = "https://login.salesforce.com/services/oauth2/token"
	}
}
