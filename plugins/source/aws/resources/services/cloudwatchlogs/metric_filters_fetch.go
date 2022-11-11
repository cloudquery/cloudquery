package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCloudwatchlogsMetricFilters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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

func resolveMetricFilterArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "logs",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "metric_filter/" + aws.ToString(resource.Item.(types.MetricFilter).FilterName),
	}
	return resource.Set(c.Name, a.String())
}
