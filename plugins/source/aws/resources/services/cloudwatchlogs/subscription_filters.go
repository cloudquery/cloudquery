package cloudwatchlogs

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func subscriptionFilters() *schema.Table {
	tableName := "aws_cloudwatchlogs_log_group_subscription_filters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_SubscriptionFilter.html`,
		Resolver:    fetchCloudwatchlogsSubscriptionFilters,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "logs"),
		Transform:   transformers.TransformWithStruct(&types.SubscriptionFilter{}, transformers.WithPrimaryKeys("FilterName", "CreationTime")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "log_group_arn",
				Description: "The Amazon Resource Name (ARN) of the log group.",
				Type:        arrow.BinaryTypes.String,
				Resolver:    schema.ParentColumnResolver("arn"),
				PrimaryKey:  true,
			},
		},
	}
}
func fetchCloudwatchlogsSubscriptionFilters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := cloudwatchlogs.DescribeSubscriptionFiltersInput{
		LogGroupName: parent.Item.(types.LogGroup).LogGroupName,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudwatchlogs
	paginator := cloudwatchlogs.NewDescribeSubscriptionFiltersPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cloudwatchlogs.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.SubscriptionFilters
	}
	return nil
}
