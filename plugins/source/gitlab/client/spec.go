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
	if s.BaseURL == "" {
		return errors.New("missing GitLab base URL in configuration file")
	}
	if s.BaseURL == "gitlab.com" {
		return errors.New("base URL cannot be `gitlab.com` this plugin only supports self hosted instances")
	}
	return nil
}
