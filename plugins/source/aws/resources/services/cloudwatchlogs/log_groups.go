package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func LogGroups() *schema.Table {
	tableName := "aws_cloudwatchlogs_log_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LogGroup.html`,
		Resolver:    fetchCloudwatchlogsLogGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "logs"),
		Transform:   client.TransformWithStruct(&types.LogGroup{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveLogGroupTags,
			},
		},
		Relations: []*schema.Table{
			subscriptionFilters(),
		},
	}
}

func fetchCloudwatchlogsLogGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config cloudwatchlogs.DescribeLogGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudwatchlogs
	paginator := cloudwatchlogs.NewDescribeLogGroupsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.LogGroups
	}
	return nil
}

func resolveLogGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	lg := resource.Item.(types.LogGroup)
	svc := meta.(*client.Client).Services().Cloudwatchlogs
	out, err := svc.ListTagsLogGroup(ctx, &cloudwatchlogs.ListTagsLogGroupInput{LogGroupName: lg.LogGroupName})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out.Tags)
}
