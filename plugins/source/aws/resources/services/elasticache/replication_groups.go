package elasticache

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ReplicationGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_replication_groups",
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ReplicationGroup.html`,
		Resolver:    fetchElasticacheReplicationGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticache"),
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
