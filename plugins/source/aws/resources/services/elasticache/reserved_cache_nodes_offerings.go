package elasticache

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ReservedCacheNodesOfferings() *schema.Table {
	tableName := "aws_elasticache_reserved_cache_nodes_offerings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ReservedCacheNodesOffering.html`,
		Resolver:    fetchElasticacheReservedCacheNodesOfferings,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticache"),
		Transform:   transformers.TransformWithStruct(&types.ReservedCacheNodesOffering{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveCacheNodesOfferingArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
