package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func MetricFilters() *schema.Table {
	tableName := "aws_cloudwatchlogs_metric_filters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_MetricFilter.html`,
		Resolver:    fetchCloudwatchlogsMetricFilters,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "logs"),
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
	paginator := cloudwatchlogs.NewDescribeMetricFiltersPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cloudwatchlogs.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- page.MetricFilters
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
