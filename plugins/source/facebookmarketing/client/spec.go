package client

import "fmt"

type Spec struct {
	AdAccountId string `json:"ad_account_id"`
	AccessToken string `json:"access_token"`
}

func (s *Spec) Validate() error {
	if s.AdAccountId == "" {
		return fmt.Errorf("ad_account_id is required")
	}

	if s.AccessToken == "" {
		return fmt.Errorf("access_token is required")
	}

	return nil
}
