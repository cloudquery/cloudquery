package cloudwatchlogs

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func LogGroups() *schema.Table {
	tableName := "aws_cloudwatchlogs_log_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LogGroup.html`,
		Resolver:    fetchCloudwatchlogsLogGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "logs"),
		Transform:   transformers.TransformWithStruct(&types.LogGroup{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Arn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveLogGroupTags,
			},
		},
		Relations: []*schema.Table{
			subscriptionFilters(),
			dataProtectionPolicy(),
		},
	}
}

func fetchCloudwatchlogsLogGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config cloudwatchlogs.DescribeLogGroupsInput
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudwatchlogs
	paginator := cloudwatchlogs.NewDescribeLogGroupsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cloudwatchlogs.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.LogGroups
	}
	return nil
}

func resolveLogGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	lg := resource.Item.(types.LogGroup)
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudwatchlogs
	out, err := svc.ListTagsLogGroup(ctx, &cloudwatchlogs.ListTagsLogGroupInput{LogGroupName: lg.LogGroupName}, func(options *cloudwatchlogs.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out.Tags)
}
