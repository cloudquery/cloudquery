package client

import "errors"

type Spec struct {
	Token   string `json:"access_token,omitempty"`
	BaseURL string `json:"base_url,omitempty"`
}

func (s Spec) Validate() error {
	gitlabToken := s.Token
	if gitlabToken == "" {
		return errors.New("missing GitLab API token in configuration file")
	}
	return nil
}
