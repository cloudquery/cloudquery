package docdb

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func EventSubscriptions() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_event_subscriptions",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_EventSubscription.html`,
		Resolver:    fetchDocdbEventSubscriptions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("docdb"),
		Transform:   transformers.TransformWithStruct(&types.EventSubscription{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
		},
	}
}
