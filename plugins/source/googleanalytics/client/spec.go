package client

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/invopop/jsonschema"
)

// CloudQuery Google Analytics source plugin confugiration spec.
type Spec struct {
	// A Google Analytics GA4 [property](https://support.google.com/analytics/answer/9304153#property) identifier whose events are tracked.
	// To learn more, see where to [find your Property ID](https://developers.google.com/analytics/devguides/reporting/data/v1/property-id).
	//
	// Supported formats:
	//
	// - A plain property ID (example: `1234`)
	//
	// - Prefixed with `properties/` (example: `properties/1234`)
	PropertyID string `json:"property_id,omitempty" jsonschema:"required,minLength=1"`

	// Reports to be fetched from Google Analytics.
	Reports []Report `json:"reports,omitempty"`

	// A date in `YYYY-MM-DD` format (example: `2023-05-15`).
	// If not specified, the start date will be the one that is 7 days prior to the sync start date.
	StartDate string `json:"start_date,omitempty" jsonschema:"format=date,default=now-168h"`

	// OAuth spec for authorization in Google Analytics.
	OAuth *OAuthSpec `json:"oauth,omitempty"`

	// The best effort maximum number of Go routines to use.
	// Lower this number to reduce memory usage.
	Concurrency int `json:"concurrency,omitempty" jsonschema:"minimum=1,default=10000"`
}

func (s *Spec) SetDefaults() {
	if len(s.StartDate) == 0 {
		// date 7 days prior
		s.StartDate = time.Now().UTC().Add(-7 * 24 * time.Hour).Format(time.DateOnly)
	}
	if s.Concurrency <= 0 {
		const defaultConcurrency = 10000
		s.Concurrency = defaultConcurrency
	}
}

func (Spec) JSONSchemaExtend(sc *jsonschema.Schema) {
}

func (s *Spec) Validate() error {
	if len(s.PropertyID) == 0 {
		return fmt.Errorf(`required field "property_id" is missing`)
	}
	if !strings.HasPrefix(s.PropertyID, "properties/") {
		s.PropertyID = "properties/" + s.PropertyID // required for SDK
	}

	_, err := time.Parse(time.DateOnly, s.StartDate)
	if err != nil {
		return fmt.Errorf(`"start_date" has to be in %q format, got %q: %w`, time.DateOnly, s.StartDate, err)
	}

	saw := make(map[string]struct{})
	for _, r := range s.Reports {
		r.normalize()
		if err := r.validate(); err != nil {
			return fmt.Errorf("failed to validate report %q: %w", r.Name, err)
		}
		if _, ok := saw[r.Name]; ok {
			return fmt.Errorf("report with name %q already present", r.Name)
		}
		saw[r.Name] = struct{}{}
	}

	return s.OAuth.validate()
}

//go:embed schema.json
var JSONSchema string
