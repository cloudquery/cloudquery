package costexplorer

import (
	"context"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ThirtyDayCost() *schema.Table {
	tableName := "aws_costexplorer_cost_30d"
	return &schema.Table{
		Name:     tableName,
		Resolver: fetchCost,
		Title:    "AWS Cost Explorer costs for the last 30 days",
		Description: `https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetCostAndUsage.html
To sync this table you must set the 'use_paid_apis' option to 'true' in the AWS provider configuration. `,
		Transform: transformers.TransformWithStruct(&types.ResultByTime{}),
		Multiplex: client.AccountMultiplex(tableName),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:        "start_date",
				Description: `The start date covered by the forecast.`,
				Type:        arrow.BinaryTypes.String,
				Resolver:    schema.PathResolver("TimePeriod.Start"),
				PrimaryKey:  true,
			},
			{
				Name:        "end_date",
				Description: `The end date covered by the forecast.`,
				Type:        arrow.BinaryTypes.String,
				Resolver:    schema.PathResolver("TimePeriod.End"),
				PrimaryKey:  true,
			},
		},
	}
}

func fetchCost(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	if !cl.Spec.UsePaidAPIs {
		cl.Logger().Info().Msg("skipping `aws_costexplorer_cost_current_month` because `use_paid_apis` is set to false")
		return nil
	}
	svc := cl.Services().Costexplorer
	// Only use a single `time.Now()` call to ensure that the start and end dates are the same.
	now := time.Now()

	input := costexplorer.GetCostAndUsageInput{
		Granularity: types.GranularityDaily,
		Metrics: []string{
			"AmortizedCost",
			"BlendedCost",
			"NetAmortizedCost",
			"NetUnblendedCost",
			"NormalizedUsageAmount",
			"UnblendedCost",
		},
		TimePeriod: &types.DateInterval{
			Start: aws.String(now.AddDate(0, 0, -30).Format("2006-01-02")),
			End:   aws.String(now.AddDate(0, 0, 1).Format("2006-01-02")),
		},
	}
	for {
		resp, err := svc.GetCostAndUsage(ctx, &input, func(options *costexplorer.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- resp.ResultsByTime
		if aws.ToString(resp.NextPageToken) == "" {
			break
		}
	}

	return nil
}
