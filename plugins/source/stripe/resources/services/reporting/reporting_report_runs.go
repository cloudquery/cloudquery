package reporting

import (
	"context"

	"fmt"
	"strconv"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func ReportingReportRuns() *schema.Table {
	return &schema.Table{
		Name:        "stripe_reporting_report_runs",
		Description: `https://stripe.com/docs/api/reporting_report_runs`,
		Transform:   transformers.TransformWithStruct(&stripe.ReportingReportRun{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
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
			{
				Name:     "created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Created"),
				CreationOptions: schema.ColumnCreationOptions{
					IncrementalKey: true,
				},
			},
		},
		IsIncremental: true,
	}
}

func fetchReportingReportRuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.ReportingReportRunListParams{}

	const key = "reporting_report_runs"

	if cl.Backend != nil {
		value, err := cl.Backend.Get(ctx, key, cl.ID())
		if err != nil {
			return fmt.Errorf("failed to retrieve state from backend: %w", err)
		}
		if value != "" {
			vi, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("retrieved invalid state value: %q %w", value, err)
			}
			lp.Created = &vi
		}
	}

	it := cl.Services.ReportingReportRuns.List(lp)
	for it.Next() {
		data := it.ReportingReportRun()
		lp.Created = client.MaxInt64(lp.Created, &data.Created)
		res <- data
	}

	err := it.Err()
	if cl.Backend != nil && err == nil && lp.Created != nil {
		return cl.Backend.Set(ctx, key, cl.ID(), strconv.FormatInt(*lp.Created, 10))
	}
	return err
}
