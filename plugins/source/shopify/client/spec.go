package client

import (
	"errors"
	"strings"
)

type Spec struct {
	// Either
	AccessToken string `json:"access_token,omitempty"`
	// or
	APIKey    string `json:"api_key,omitempty"`
	APISecret string `json:"api_secret,omitempty"`

	// Required
	ShopURL string `json:"shop_url"`

	// Optional
	Timeout     int64 `json:"timeout_secs,omitempty"`
	MaxRetries  int64 `json:"max_retries,omitempty"`
	PageSize    int64 `json:"page_size,omitempty"`
	Concurrency int   `json:"concurrency,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Timeout < 1 {
		s.Timeout = 10
	}
	if s.MaxRetries < 1 {
		s.MaxRetries = 30
	}
	if s.PageSize < 1 {
		s.PageSize = 50
	}
	if s.Concurrency < 1 {
		s.Concurrency = 1000
	}
}

func (s Spec) Validate() error {
	if s.AccessToken == "" && (s.APIKey == "" || s.APISecret == "") {
		return errors.New("no credentials provided")
	}
	if s.ShopURL == "" {
		return errors.New("no shop url provided")
	}
	if !strings.HasSuffix(s.ShopURL, ".myshopify.com") {
		return errors.New("shop url should end with .myshopify.com, as in https://shop_name.myshopify.com")
	}

	return nil
}
