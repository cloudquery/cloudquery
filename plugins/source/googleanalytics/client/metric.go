package client

import (
	"fmt"

	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
)

type Metric analyticsdata.Metric

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
