package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func MetricFilters() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudwatchlogs_metric_filters",
		Description: `https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_MetricFilter.html`,
		Resolver:    fetchCloudwatchlogsMetricFilters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("logs"),
		Transform:   transformers.TransformWithStruct(&types.MetricFilter{}, transformers.WithPrimaryKeys("FilterName")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "log_group_arn",
				Type:     schema.TypeString,
				Resolver: resolveMetricFilterLogGroupArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchCloudwatchlogsMetricFilters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config cloudwatchlogs.DescribeMetricFiltersInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudwatchlogs
	for {
		response, err := svc.DescribeMetricFilters(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.MetricFilters
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func resolveMetricFilterLogGroupArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.MetricFilter)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "logs",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "log-group:" + aws.ToString(r.LogGroupName),
	}
	return resource.Set(c.Name, a.String())
}
