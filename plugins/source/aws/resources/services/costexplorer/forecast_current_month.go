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
	tableName := "awscost_costexplorer_forecast_current_month"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/xray/latest/api/API_Group.html`,
		Resolver:    fetchForecast,
		Transform:   transformers.TransformWithStruct(&types.ForecastResult{}),
		Multiplex:   client.AccountMultiplex(tableName),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
		},
	}
}

func fetchForecast(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	if !cl.Spec.SkipNonFreeAPIs {
		cl.Logger().Info().Msg("skipping `awscost_costexplorer_forecast_current_month` because `use_non_free_apis` is set to false")
		return nil
	}

	svc := cl.Services().Costexplorer
	input := costexplorer.GetCostForecastInput{
		Granularity: types.GranularityDaily,
		Metric:      types.MetricBlendedCost,
		TimePeriod: &types.DateInterval{
			Start: aws.String(string(time.Now().Format("2006-01-02"))),
			End:   aws.String(string(endOfMonth(time.Now()).Format("2006-01-02"))),
		},
	}

	resp, err := svc.GetCostForecast(ctx, &input)
	if err != nil {
		return err
	}
	res <- resp.ForecastResultsByTime

	return nil
}
