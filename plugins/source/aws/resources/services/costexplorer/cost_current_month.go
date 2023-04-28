package costexplorer

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func CurrentMonthCost() *schema.Table {
	tableName := "aws_costexplorer_cost_current_month"
	return &schema.Table{
		Name:     tableName,
		Resolver: fetchCost,
		Description: `https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetCostAndUsage.html
To sync this table you must set the 'use_paid_apis' option to 'true' in the AWS provider configuration. `,
		Transform: transformers.TransformWithStruct(&types.ResultByTime{}),
		Multiplex: client.AccountMultiplex(tableName),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:        "start_date",
				Description: `The start date covered by the forecast.`,
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TimePeriod.Start"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "end_date",
				Description: `The end date covered by the forecast.`,
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TimePeriod.End"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
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
			Start: aws.String(beginningOfMonth(time.Now()).Format("2006-01-02")),
			End:   aws.String(time.Now().AddDate(0, 0, 1).Format("2006-01-02")),
		},
	}
	for {
		resp, err := svc.GetCostAndUsage(ctx, &input)
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
