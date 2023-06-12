package cloudwatch

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func MetricStatisticsCustom() *schema.Table {
	tableName := "aws_cloudwatch_metric_statistics_custom"
	return &schema.Table{
		Name:  tableName,
		Title: "Cloudwatch Metric Statistics custom query",
		Description: `https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricStatistics.html
To sync this table you must set the 'use_paid_apis' option to 'true' and set the relevant 'table_options' entry in the AWS provider configuration.
`,
		Resolver:  fetchCloudwatchMetricStatisticsCustom,
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "monitoring"),
		Transform: transformers.TransformWithStruct(&statOutput{}, transformers.WithSkipFields("ResultMetadata"), transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchCloudwatchMetricStatisticsCustom(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	var allConfigs []tableoptions.CloudwatchGetMetricStatisticsInput
	if cl.Spec.TableOptions.CloudwatchCustomMetricStats != nil {
		allConfigs = cl.Spec.TableOptions.CloudwatchCustomMetricStats.GetMetricStatisticsOpts
	}

	if len(allConfigs) > 0 && !cl.Spec.UsePaidAPIs {
		return client.ErrPaidAPIsNotEnabled
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
