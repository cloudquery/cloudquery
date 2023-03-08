package elasticache

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Events() *schema.Table {
	tableName := "aws_elasticache_events"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_Event.html`,
		Resolver:    fetchElasticacheEvents,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticache"),
		Transform:   transformers.TransformWithStruct(&types.Event{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "_event_hash",
				Type:     schema.TypeString,
				Resolver: client.ResolveObjectHash,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
