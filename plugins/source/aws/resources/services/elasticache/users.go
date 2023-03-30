package elasticache

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Users() *schema.Table {
	tableName := "aws_elasticache_users"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_User.html`,
		Resolver:    fetchElasticacheUsers,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticache"),
		Transform:   transformers.TransformWithStruct(&types.User{}),
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
