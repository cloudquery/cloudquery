package maintenance_windows

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func MaintenanceWindows() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_maintenance_windows",
		Description: `https://developer.pagerduty.com/api-reference/4c0936c241cbb-list-maintenance-windows`,
		Resolver:    fetchMaintenanceWindows,
		Transform:   transformers.TransformWithStruct(&pagerduty.MaintenanceWindow{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL")),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "html_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "start_time",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: schema.PathResolver("StartTime"),
			},
			{
				Name:     "end_time",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: schema.PathResolver("EndTime"),
			},
		},
	}
}
