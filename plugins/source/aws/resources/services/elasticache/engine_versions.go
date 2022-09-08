package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EngineVersions() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_engine_versions",
		Description: "Provides all of the details about a particular cache engine version.",
		Resolver:    fetchElasticacheEngineVersions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticache"),
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "region",
				Description:     "The AWS Region of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "cache_engine_description",
				Description: "The description of the cache engine.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cache_engine_version_description",
				Description: "The description of the cache engine version.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cache_parameter_group_family",
				Description: "The name of the cache parameter group family associated with this cache engine. Valid values are: memcached1.4 \\| memcached1.5 \\| memcached1.6 \\| redis2.6 \\| redis2.8 \\| redis3.2 \\| redis4.0 \\| redis5.0 \\| redis6.x",
				Type:        schema.TypeString,
			},
			{
				Name:            "engine",
				Description:     "The name of the cache engine.",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "engine_version",
				Description:     "The version number of the cache engine.",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchElasticacheEngineVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := elasticache.NewDescribeCacheEngineVersionsPaginator(meta.(*client.Client).Services().ElastiCache, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- v.CacheEngineVersions
	}
	return nil
}
