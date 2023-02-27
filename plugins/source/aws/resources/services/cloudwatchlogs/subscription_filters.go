package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func subscriptionFilters() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudwatchlogs_log_group_subscription_filters",
		Description: `https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_SubscriptionFilter.html`,
		Resolver:    fetchCloudwatchlogsSubscriptionFilters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("logs"),
		Transform:   transformers.TransformWithStruct(&types.SubscriptionFilter{}, transformers.WithPrimaryKeys("FilterName", "CreationTime")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "log_group_arn",
				Description: "The Amazon Resource Name (ARN) of the log group.",
				Type:        schema.TypeString,
				Resolver:    schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
func fetchCloudwatchlogsSubscriptionFilters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := cloudwatchlogs.DescribeSubscriptionFiltersInput{
		LogGroupName: parent.Item.(types.LogGroup).LogGroupName,
	}
	c := meta.(*client.Client)
	svc := c.Services().Cloudwatchlogs
	paginator := cloudwatchlogs.NewDescribeSubscriptionFiltersPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.SubscriptionFilters
	}
	return nil
}
