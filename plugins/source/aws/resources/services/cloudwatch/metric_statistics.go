package cloudwatch

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

type statOutput struct {
	*cloudwatch.GetMetricStatisticsOutput
	InputJSON tableoptions.CloudwatchGetMetricStatisticsInput `json:"input_json"`
}

func metricStatistics() *schema.Table {
	tableName := "aws_cloudwatch_metric_stats"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html`,
		Resolver:    fetchCloudwatchMetricStats,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "monitoring"),
		Transform:   transformers.TransformWithStruct(&statOutput{}, transformers.WithSkipFields("ResultMetadata"), transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchCloudwatchMetricStats(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	item := parent.Item.(metricOutput)

	svc := cl.Services().Cloudwatch
	for _, input := range item.getStatsInputs {
		input := input

		input.Dimensions = item.Dimensions
		input.Namespace = item.Namespace
		input.MetricName = item.MetricName

		data, err := svc.GetMetricStatistics(ctx, &input.GetMetricStatisticsInput, func(options *cloudwatch.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- statOutput{
			GetMetricStatisticsOutput: data,
			InputJSON:                 input,
		}
	}
	return nil
}
