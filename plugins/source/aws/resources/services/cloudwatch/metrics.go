package cloudwatch

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client/tableoptions"
	sdkTypes "github.com/cloudquery/plugin-sdk/v3/types"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Metrics() *schema.Table {
	tableName := "aws_cloudwatch_metrics_custom"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Metric.html`,
		Resolver:    fetchCloudwatchMetrics,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "logs"),
		Transform:   transformers.TransformWithStruct(&types.Metric{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "dimensions",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveCloudwatchMetricDimensions,
			},
		},
	}
}

func fetchCloudwatchMetrics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	var allConfigs []tableoptions.CustomCloudwatchListMetricsInput
	if cl.Spec.TableOptions.CloudwatchEvents != nil {
		allConfigs = cl.Spec.TableOptions.CloudwatchEvents.ListMetricsOpts
	}

	if len(allConfigs) > 0 && !cl.Spec.UsePaidAPIs {
		cl.Logger().Info().Msg("skipping `aws_cloudwatch_metrics_custom` because `use_paid_apis` is set to false")
		return nil
	}

	svc := cl.Services().Cloudwatch
	for _, input := range allConfigs {
		paginator := cloudwatch.NewListMetricsPaginator(svc, &input.ListMetricsInput)
		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx, func(options *cloudwatch.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- page.Metrics
		}
	}
	return nil
}

func resolveCloudwatchMetricDimensions(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	item := resource.Item.(types.Metric)
	dimensions := make(map[string]*string)
	for _, d := range item.Dimensions {
		dimensions[*d.Name] = d.Value
	}
	return resource.Set(c.Name, dimensions)
}
