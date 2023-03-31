package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ParameterGroups() *schema.Table {
	tableName := "aws_elasticache_parameter_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheParameterGroup.html`,
		Resolver:    fetchElasticacheParameterGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticache"),
		Transform:   transformers.TransformWithStruct(&types.CacheParameterGroup{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchElasticacheParameterGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	awsProviderClient := meta.(*client.Client)
	svc := awsProviderClient.Services().Elasticache

	var describeCacheParameterGroupsInput elasticache.DescribeCacheParameterGroupsInput
	paginator := elasticache.NewDescribeCacheParameterGroupsPaginator(svc, &describeCacheParameterGroupsInput)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.CacheParameterGroups
	}
	return nil
}

func fetchElasticacheGlobalReplicationGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	paginator := elasticache.NewDescribeGlobalReplicationGroupsPaginator(meta.(*client.Client).Services().Elasticache, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- v.GlobalReplicationGroups
	}
	return nil
}
