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
	InputJSON tableoptions.CustomCloudwatchGetMetricStatisticsInput `json:"input_json"`
}

func MetricStatistics() *schema.Table {
	tableName := "aws_cloudwatch_metric_stats_custom"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html`,
		Resolver:    fetchCloudwatchMetricStats,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "logs"),
		Transform:   transformers.TransformWithStruct(&statOutput{}, transformers.WithSkipFields("ResultMetadata"), transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchCloudwatchMetricStats(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	var allConfigs []tableoptions.CustomCloudwatchGetMetricStatisticsInput
	if cl.Spec.TableOptions.CloudwatchMetricStats != nil {
		allConfigs = cl.Spec.TableOptions.CloudwatchMetricStats.GetMetricStatisticsOpts
	}

	if len(allConfigs) > 0 && !cl.Spec.UsePaidAPIs {
		cl.Logger().Info().Msg("skipping `aws_cloudwatch_metric_stats_custom` because `use_paid_apis` is set to false")
		return nil
	}

	svc := cl.Services().Cloudwatch
	for _, input := range allConfigs {
		input := input
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
