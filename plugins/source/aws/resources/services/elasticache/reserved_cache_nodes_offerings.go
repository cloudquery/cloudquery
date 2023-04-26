package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
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

func fetchElasticacheReservedCacheNodesOfferings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	paginator := elasticache.NewDescribeReservedCacheNodesOfferingsPaginator(meta.(*client.Client).Services().Elasticache, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- v.ReservedCacheNodesOfferings
	}
	return nil
}

func resolveCacheNodesOfferingArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.ReservedCacheNodesOffering)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "elasticache",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "elasticache/" + aws.ToString(item.ReservedCacheNodesOfferingId),
	}
	return resource.Set(c.Name, a.String())
}
