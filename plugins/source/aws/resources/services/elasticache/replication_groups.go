package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func ReplicationGroups() *schema.Table {
	tableName := "aws_elasticache_replication_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ReplicationGroup.html`,
		Resolver:    fetchElasticacheReplicationGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticache"),
		Transform:   transformers.TransformWithStruct(&types.ReplicationGroup{}),
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
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveElasticacheReplicationGroupTags,
			},
		},
	}
}

func fetchElasticacheReplicationGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	paginator := elasticache.NewDescribeReplicationGroupsPaginator(meta.(*client.Client).Services().Elasticache, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- v.ReplicationGroups
	}
	return nil
}

func resolveElasticacheReplicationGroupTags(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	svc := meta.(*client.Client).Services().Elasticache
	tags, err := svc.ListTagsForResource(ctx, &elasticache.ListTagsForResourceInput{ResourceName: r.Item.(types.ReplicationGroup).ARN})
	if err != nil {
		return err
	}
	return r.Set(c.Name, client.TagsToMap(tags.TagList))
}
