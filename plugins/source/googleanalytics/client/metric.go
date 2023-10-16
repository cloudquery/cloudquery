package client

import (
	"fmt"

	"github.com/invopop/jsonschema"
	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
)

type Metric analyticsdata.Metric

func (Metric) JSONSchemaExtend(sc *jsonschema.Schema) {
	name := sc.Properties.Value("name")
	one := uint64(1)
	name.MinLength = &one

	sc.Required = append(sc.Required, "name")
}

func (m *Metric) validate() error {
	switch {
	case m == nil:
		return fmt.Errorf("empty metric spec")
	case len(m.Name) == 0:
		return fmt.Errorf("empty metric name")
	default:
		return nil
	}
}

func (m *Metric) toGA() *analyticsdata.Metric {
	res := analyticsdata.Metric(*m)
	return &res
}
