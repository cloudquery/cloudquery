package elasticache

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Snapshots() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_snapshots",
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_Snapshot.html`,
		Resolver:    fetchElasticacheSnapshots,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticache"),
		Transform:   transformers.TransformWithStruct(&types.Snapshot{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
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
