package reporting

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func ReportingReportRuns() *schema.Table {
	return &schema.Table{
		Name:        "stripe_reporting_report_runs",
		Description: `https://stripe.com/docs/api/reporting_report_runs`,
		Transform:   transformers.TransformWithStruct(&stripe.ReportingReportRun{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchReportingReportRuns,

		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchReportingReportRuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.ReportingReportRuns.List(&stripe.ReportingReportRunListParams{})
	for it.Next() {
		res <- it.ReportingReportRun()
	}
	return it.Err()
}
