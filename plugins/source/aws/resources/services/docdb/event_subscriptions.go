package docdb

import (
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func EventSubscriptions() *schema.Table {
	tableName := "aws_docdb_event_subscriptions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_EventSubscription.html`,
		Resolver:    fetchDocdbEventSubscriptions,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "docdb"),
		Transform:   transformers.TransformWithStruct(&types.EventSubscription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}
