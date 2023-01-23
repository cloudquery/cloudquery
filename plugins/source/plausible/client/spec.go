package client

import (
	"fmt"

	"golang.org/x/exp/slices"
)

var metrics = []string{"visitors", "pageviews", "bounce_rate", "visit_duration", "visits"}

const baseURL = "https://plausible.io"

type Spec struct {
	SiteId   string   `json:"site_id"`
	ApiKey   string   `json:"api_key"`
	BaseURL  string   `json:"base_url"`
	Period   string   `json:"period"`
	Filters  string   `json:"filters"`
	Metrics  []string `json:"metrics"`
	Interval string   `json:"interval"`
}

func (s *Spec) Validate() error {
	if s.SiteId == "" {
		return fmt.Errorf("site_id is required")
	}
	if s.ApiKey == "" {
		return fmt.Errorf("api_key is required")
	}
	if len(s.Metrics) != 0 {
		for _, m := range s.Metrics {
			if slices.Index(metrics, m) == -1 {
				return fmt.Errorf("invalid metric: %s. must be one of %v", m, metrics)
			}
		}
	}
	return nil
}

func (s *Spec) SetDefaults() {
	if s.BaseURL == "" {
		s.BaseURL = baseURL
	}
	if s.Period == "" {
		s.Period = "30d"
	}
	if s.Metrics == nil {
		s.Metrics = metrics
	}
	if s.Interval == "" {
		s.Interval = "date"
	}
}
