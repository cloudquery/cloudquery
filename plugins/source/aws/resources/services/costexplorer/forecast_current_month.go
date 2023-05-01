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

func CurrentMonthForecast() *schema.Table {
	tableName := "aws_costexplorer_forecast_current_month"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_GetCostForecast.html
To sync this table you must set the 'use_paid_apis' option to 'true' in the AWS provider configuration. `,
		Resolver:  fetchForecast,
		Transform: transformers.TransformWithStruct(&types.ForecastResult{}),
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

func fetchForecast(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	if !cl.Spec.UsePaidAPIs {
		cl.Logger().Info().Msg("skipping `aws_costexplorer_forecast_current_month` because `use_paid_apis` is set to false")
		return nil
	}
	now := time.Now()
	svc := cl.Services().Costexplorer

	input := costexplorer.GetCostForecastInput{
		Granularity: types.GranularityDaily,
		Metric:      types.MetricBlendedCost,
		TimePeriod: &types.DateInterval{
			Start: aws.String(now.Format("2006-01-02")),
			End:   aws.String(forecastEndOfMonth(now).Format("2006-01-02")),
		},
	}

	resp, err := svc.GetCostForecast(ctx, &input)
	if err != nil {
		return err
	}
	res <- resp.ForecastResultsByTime
	return nil
}
