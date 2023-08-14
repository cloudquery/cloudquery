package client

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const defaultPeriod = "30d"

var reValidPeriod = regexp.MustCompile(`^(\d)+([dw])$`)

type Spec struct {
	// APIKey required to access Snyk API
	APIKey string `json:"api_key,omitempty"`

	// Organizations is a list of organizations to fetch information from.
	// By default, will fetch from all organizations available for user.
	Organizations []string `json:"organizations,omitempty"`

	// EndpointURL is an optional parameter to override the API URL for snyk.Client.
	// It defaults to https://api.snyk.io/api/
	EndpointURL string `json:"endpoint_url,omitempty"`

	// Retries is an optional parameter to override the default number of retries for retryable requests.
	Retries int `json:"retries,omitempty"`

	// RetryDelaySeconds is an optional parameter to override the default backoff time for retryable requests.
	RetryDelaySeconds int `json:"retry_delay_seconds,omitempty"`

	TableOptions TableOptions `json:"table_options,omitempty"`

	Concurrency int `json:"concurrency,omitempty"`
}

type TableOptions struct {
	SnykReportingIssues SnykReportingIssuesOptions `json:"snyk_reporting_issues"`
}

func (t TableOptions) Validate() error {
	if t.SnykReportingIssues.From != "" {
		_, err := time.Parse("2006-01-02", t.SnykReportingIssues.From)
		if err != nil {
			return fmt.Errorf("invalid from date: %w", err)
		}
	}
	if t.SnykReportingIssues.To != "" {
		_, err := time.Parse("2006-01-02", t.SnykReportingIssues.To)
		if err != nil {
			return fmt.Errorf("invalid to date: %w", err)
		}
	}
	if t.SnykReportingIssues.Period != "" {
		if !reValidPeriod.MatchString(t.SnykReportingIssues.Period) {
			return fmt.Errorf("invalid period: %s", t.SnykReportingIssues.Period)
		}
	}
	if t.SnykReportingIssues.To != "" && t.SnykReportingIssues.From == "" {
		return errors.New("cannot use to without from")
	}
	if t.SnykReportingIssues.Period != "" && (t.SnykReportingIssues.From != "" || t.SnykReportingIssues.To != "") {
		return errors.New("cannot use period with to/from")
	}
	return nil
}

// SnykReportingIssuesOptions accepts these combinations:
// - `from` + `to`
// - `from` (default `to` is now)
// - `period` (relative to now)
// Other combinations should fail validation.
type SnykReportingIssuesOptions struct {
	From   string `json:"from"`   // e.g. 2020-01-01
	To     string `json:"to"`     // e.g. 2020-01-01
	Period string `json:"period"` // e.g. 1d, 365d
}

func (s SnykReportingIssuesOptions) FromTime() time.Time {
	if s.From != "" {
		t, _ := time.Parse("2006-01-02", s.From)
		return t
	}
	period := defaultPeriod
	if s.Period != "" {
		period = s.Period
	}
	return time.Now().Add(-parseDuration(period)).Truncate(24 * time.Hour)
}

func (s SnykReportingIssuesOptions) ToTime() time.Time {
	if s.To != "" {
		t, _ := time.Parse("2006-01-02", s.To)
		return t
	}
	return time.Now().Truncate(24 * time.Hour)
}

func parseDuration(s string) time.Duration {
	matches := reValidPeriod.FindAllStringSubmatch(s, -1)
	if len(matches) == 0 {
		return 0
	}
	switch matches[0][2] {
	case "d":
		return time.Duration(24*int64(time.Hour)) * time.Duration(mustParseInt(matches[0][1]))
	case "w":
		return time.Duration(24*7*int64(time.Hour)) * time.Duration(mustParseInt(matches[0][1]))
	}
	return 0
}

func mustParseInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (s *Spec) Validate() error {
	if len(s.APIKey) == 0 {
		return fmt.Errorf("missing API key")
	}
	return s.TableOptions.Validate()
}

func (s *Spec) SetDefaults() {
	if s.Concurrency < 1 {
		s.Concurrency = 10000
	}
}
