package elasticache

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ServiceUpdates() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_service_updates",
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ServiceUpdate.html`,
		Resolver:    fetchElasticacheServiceUpdates,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticache"),
		Transform:   transformers.TransformWithStruct(&types.ServiceUpdate{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveServiceUpdateArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
