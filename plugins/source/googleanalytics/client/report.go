package client

import (
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/invopop/jsonschema"
	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
)

type Report struct {
	Name          string    `json:"name" jsonschema:"required,minLength=1"`
	Dimensions    []string  `json:"dimensions,omitempty" jsonschema:"maxItems=9,minLength=1"`
	Metrics       []*Metric `json:"metrics" jsonschema:"required,minItems=1"`
	KeepEmptyRows bool      `json:"keep_empty_rows,omitempty"`
}

var csr = caser.New()

func (r *Report) normalize() {
	if r == nil {
		return
	}

	r.Name = csr.ToSnake(strings.ReplaceAll(r.Name, " ", "_"))
}

func (Report) JSONSchemaExtend(sc *jsonschema.Schema) {
	metrics := sc.Properties.Value("metrics").OneOf[0] // 0 - spec, 1 - null
	sc.Properties.Set("metrics", metrics)
}

func (r *Report) validate() error {
	switch {
	case r == nil:
		return fmt.Errorf("empty report spec")
	case len(r.Name) == 0:
		return fmt.Errorf("empty report name")
	case len(r.Dimensions) > 9:
		return fmt.Errorf("report %q: requested %d dimensions (can have up to 9 only)", r.Name, len(r.Dimensions))
	case len(r.Metrics) == 0:
		return fmt.Errorf("empty report metrics")
	}

	for _, m := range r.Metrics {
		if err := m.validate(); err != nil {
			return err
		}
	}

	return nil
}

func (r *Report) toGA(propertyID string) *analyticsdata.RunReportRequest {
	req := &analyticsdata.RunReportRequest{
		Property:      propertyID,
		Dimensions:    make([]*analyticsdata.Dimension, len(r.Dimensions)),
		Metrics:       make([]*analyticsdata.Metric, len(r.Metrics)),
		KeepEmptyRows: r.KeepEmptyRows,
	}

	for i, d := range r.Dimensions {
		req.Dimensions[i] = &analyticsdata.Dimension{Name: d}
	}

	for i, m := range r.Metrics {
		req.Metrics[i] = m.toGA()
	}

	return req
}

func (r *Report) table(propertyID string) *schema.Table {
	tableName := "ga_" + r.Name
	return &schema.Table{
		Name:          tableName,
		Transform:     transformers.TransformWithStruct(new(row), transformers.WithPrimaryKeys("DimensionHash")),
		Columns:       schema.ColumnList{PropertyIDColumn, DateColumn},
		Resolver:      fetch(tableName, r.toGA(propertyID)),
		IsIncremental: true,
	}
}
