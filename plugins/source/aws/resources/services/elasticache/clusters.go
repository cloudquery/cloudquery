package elasticache

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Clusters() *schema.Table {
	tableName := "aws_elasticache_clusters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheCluster.html`,
		Resolver:    fetchElasticacheClusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticache"),
		Transform:   transformers.TransformWithStruct(&types.CacheCluster{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ARN"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveClusterTags,
			},
		},
	}
}

func fetchElasticacheClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceElasticache).Elasticache
	var input elasticache.DescribeCacheClustersInput
	input.ShowCacheNodeInfo = aws.Bool(true)

	paginator := elasticache.NewDescribeCacheClustersPaginator(svc, &input)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx, func(options *elasticache.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- v.CacheClusters
	}
	return nil
}

func resolveClusterTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cluster := resource.Item.(types.CacheCluster)

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceElasticache).Elasticache
	response, err := svc.ListTagsForResource(ctx, &elasticache.ListTagsForResourceInput{
		ResourceName: cluster.ARN,
	}, func(options *elasticache.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(response.TagList))
}
