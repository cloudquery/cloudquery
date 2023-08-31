package client

import "fmt"

type Spec struct {
	ClientId       string   `json:"client_id"`
	ClientSecret   string   `json:"client_secret"`
	Username       string   `json:"username"`
	Password       string   `json:"password"`
	IncludeObjects []string `json:"include_objects"`
	ExcludeObjects []string `json:"exclude_objects"`

	Concurrency int `json:"concurrency,omitempty"`
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
	return nil
}

func (s *Spec) SetDefaults() {
	if s.IncludeObjects == nil {
		s.IncludeObjects = []string{"*"}
	}
	if s.ExcludeObjects == nil {
		s.ExcludeObjects = []string{}
	}
	if s.Concurrency < 1 {
		s.Concurrency = 10000
	}
}
