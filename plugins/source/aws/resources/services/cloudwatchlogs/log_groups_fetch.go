package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCloudwatchlogsLogGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config cloudwatchlogs.DescribeLogGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudwatchlogs
	for {
		response, err := svc.DescribeLogGroups(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.LogGroups
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveLogGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	lg := resource.Item.(types.LogGroup)
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudwatchlogs
	out, err := svc.ListTagsLogGroup(ctx, &cloudwatchlogs.ListTagsLogGroupInput{LogGroupName: lg.LogGroupName})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out.Tags)
}
