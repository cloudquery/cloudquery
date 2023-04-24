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
	tableName := "awscost_costexplorer_cost_current_month"
	return &schema.Table{
		Name:      tableName,
		Resolver:  fetchCost,
		Transform: transformers.TransformWithStruct(&types.ResultByTime{}),
		Multiplex: client.AccountMultiplex(tableName),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
		},
	}
}

func fetchCost(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
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
			Start: aws.String(string(beginningOfMonth(time.Now()).Format("2006-01-02"))),
			End:   aws.String(string(time.Now().Format("2006-01-02"))),
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
