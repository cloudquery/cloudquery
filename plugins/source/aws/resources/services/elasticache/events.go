package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
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

func fetchElasticacheEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input elasticache.DescribeEventsInput

	paginator := elasticache.NewDescribeEventsPaginator(meta.(*client.Client).Services().Elasticache, &input)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- v.Events
	}
	return nil
}
