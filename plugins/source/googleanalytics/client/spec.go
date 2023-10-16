package client

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/invopop/jsonschema"
)

type Spec struct {
	PropertyID  string     `json:"property_id,omitempty" jsonschema:"required,minLength=1"`
	StartDate   string     `json:"start_date,omitempty" jsonschema:"format=date,default=now-168h"`
	OAuth       *OAuthSpec `json:"oauth,omitempty"`
	Reports     []Report   `json:"reports,omitempty"`
	Concurrency int        `json:"concurrency,omitempty" jsonschema:"minimum=1,default=10000"`
}

func (s *Spec) setDefaults() {
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

func (s *Spec) validate() error {
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
