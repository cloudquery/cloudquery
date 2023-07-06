package elasticache

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveCacheNodesOfferingArn,
				PrimaryKey: true,
			},
		},
	}
}

func fetchElasticacheReservedCacheNodesOfferings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	paginator := elasticache.NewDescribeReservedCacheNodesOfferingsPaginator(meta.(*client.Client).Services().Elasticache, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx, func(options *elasticache.Options) {
			options.Region = cl.Region
		})
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
