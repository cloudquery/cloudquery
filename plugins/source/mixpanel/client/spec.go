package client

import (
	"errors"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/rs/zerolog"
)

type Spec struct {
	// Required fields
	Username  string `json:"username,omitempty"`
	Secret    string `json:"secret,omitempty"`
	ProjectID int64  `json:"project_id,omitempty"`

	// Optional fields
	WorkspaceID int64  `json:"workspace_id,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	Region      string `json:"region,omitempty"`
	Timeout     int64  `json:"timeout_secs,omitempty"`
	MaxRetries  int64  `json:"max_retries,omitempty"`

	// Internal fields
	region mixpanel.Region
}

const dateFormat = "2006-01-02"

func (s *Spec) SetDefaults(logger zerolog.Logger) {
	if s.StartDate == "" {
		dt := time.Now().UTC().AddDate(0, 0, -30).Format(dateFormat)
		logger.Info().Str("start_date", dt).Msg("no start date provided, defaulting to 30 days ago")
		s.StartDate = dt
	}
	if s.EndDate == "" {
		dt := time.Now().UTC().Format(dateFormat)
		logger.Info().Str("end_date", dt).Msg("no end date provided, defaulting to today")
		s.EndDate = dt
	}

	if s.Region == "" {
		s.Region = string(mixpanel.RegionUS)
	}

	if s.Timeout < 1 {
		s.Timeout = 30
	}
	if s.MaxRetries < 1 {
		s.MaxRetries = 5
	}
}

func (s *Spec) Validate() error {
	if s.Secret == "" {
		// allow empty username for the deprecated project secret method: https://developer.mixpanel.com/reference/project-secret
		return errors.New("no credentials provided")
	}
	if s.ProjectID < 1 {
		return errors.New("no project id provided")
	}

	if _, err := time.Parse(dateFormat, s.StartDate); err != nil {
		return fmt.Errorf("invalid start date format: %w", err)
	}
	if _, err := time.Parse(dateFormat, s.EndDate); err != nil {
		return fmt.Errorf("invalid end date format: %w", err)
	}

	var err error
	s.region, err = mixpanel.ParseRegion(s.Region)
	if err != nil {
		return err
	}

	return nil
}
