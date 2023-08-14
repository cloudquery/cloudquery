package reporting

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func ReportingReportTypes() *schema.Table {
	return &schema.Table{
		Name:        "stripe_reporting_report_types",
		Description: `https://stripe.com/docs/api/reporting/report_type`,
		Transform:   client.TransformWithStruct(&stripe.ReportingReportType{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchReportingReportTypes,

		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchReportingReportTypes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.ReportingReportTypeListParams{}

	it := cl.Services.ReportingReportTypes.List(lp)
	for it.Next() {
		res <- it.ReportingReportType()
	}

	return it.Err()
}
