package redshift

import (
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func EventSubscriptions() *schema.Table {
	tableName := "aws_redshift_event_subscriptions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_EventSubscription.html`,
		Resolver:    fetchRedshiftEventSubscriptions,
		Transform:   transformers.TransformWithStruct(&types.EventSubscription{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "redshift"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "arn",
				Type:        schema.TypeString,
				Resolver:    resolveEventSubscriptionARN,
				Description: `ARN of the event subscription.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
