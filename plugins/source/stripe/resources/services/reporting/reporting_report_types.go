package reporting

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func ReportingReportTypes() *schema.Table {
	return &schema.Table{
		Name:        "stripe_reporting_report_types",
		Description: `https://stripe.com/docs/api/reporting_report_types`,
		Transform:   transformers.TransformWithStruct(&stripe.ReportingReportType{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchReportingReportTypes,

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

func fetchReportingReportTypes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.ReportingReportTypeListParams{}

	it := cl.Services.ReportingReportTypes.List(lp)
	for it.Next() {
		res <- it.ReportingReportType()
	}

	return it.Err()
}
