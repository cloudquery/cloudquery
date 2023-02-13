package rds

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func EventSubscriptions() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_event_subscriptions",
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_EventSubscription.html`,
		Resolver:    fetchRdsEventSubscriptions,
		Transform:   transformers.TransformWithStruct(&types.EventSubscription{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("rds"),
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
				Resolver: resolveRDSEventSubscriptionTags,
			},
		},
	}
}
