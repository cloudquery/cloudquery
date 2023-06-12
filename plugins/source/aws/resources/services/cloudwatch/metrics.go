package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

type metricOutput struct {
	types.Metric
	InputJSON tableoptions.CloudwatchListMetricsInput `json:"input_json"`

	getStatsInputs []tableoptions.CloudwatchGetMetricStatisticsInput
}

func Metrics() *schema.Table {
	tableName := "aws_cloudwatch_metrics"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html`,
		Resolver:    fetchCloudwatchMetrics,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "monitoring"),
		Transform:   transformers.TransformWithStruct(&metricOutput{}, transformers.WithSkipFields("ResultMetadata"), transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
		Relations: []*schema.Table{
			metricStatistics(),
		},
	}
}

func fetchCloudwatchMetrics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	if len(cl.Spec.TableOptions.CloudwatchMetrics) > 0 && !cl.Spec.UsePaidAPIs {
		cl.Logger().Info().Msg("skipping `aws_cloudwatch_metrics` because `use_paid_apis` is set to false")
		return nil
	}

	svc := cl.Services().Cloudwatch
	for _, input := range cl.Spec.TableOptions.CloudwatchMetrics {
		input := input
		paginator := cloudwatch.NewListMetricsPaginator(svc, &input.ListMetricsOpts.ListMetricsInput)
		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx, func(options *cloudwatch.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			for i := range page.Metrics {
				res <- metricOutput{
					Metric:         page.Metrics[i],
					InputJSON:      input.ListMetricsOpts,
					getStatsInputs: input.GetMetricStatisticsOpts,
				}
			}
		}
	}
	return nil
}
