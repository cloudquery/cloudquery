package elasticache

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func EngineVersions() *schema.Table {
	tableName := "aws_elasticache_engine_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheEngineVersion.html`,
		Resolver:    fetchElasticacheEngineVersions,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticache"),
		Transform:   transformers.TransformWithStruct(&types.CacheEngineVersion{}),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        arrow.BinaryTypes.String,
				Resolver:    client.ResolveAWSAccount,
				Description: `The AWS Account ID of the resource.`,
				PrimaryKey:  true,
			},
			{
				Name:        "region",
				Type:        arrow.BinaryTypes.String,
				Resolver:    client.ResolveAWSRegion,
				Description: `The AWS Region of the resource.`,
				PrimaryKey:  true,
			},
			{
				Name:       "engine",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Engine"),
				PrimaryKey: true,
			},
			{
				Name:       "engine_version",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("EngineVersion"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchElasticacheEngineVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Elasticache

	paginator := elasticache.NewDescribeCacheEngineVersionsPaginator(svc, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx, func(options *elasticache.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- v.CacheEngineVersions
	}
	return nil
}
