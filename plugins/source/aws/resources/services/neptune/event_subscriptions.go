package neptune

import (
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EventSubscriptions() *schema.Table {
	tableName := "aws_neptune_event_subscriptions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/neptune/latest/userguide/api-events.html#DescribeEventSubscriptions`,
		Resolver:    fetchNeptuneEventSubscriptions,
		Transform:   client.TransformWithStruct(&types.EventSubscription{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "neptune"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventSubscriptionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveNeptuneEventSubscriptionTags,
			},
		},
	}
}
