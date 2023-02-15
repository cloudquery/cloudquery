package elasticache

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func UpdateActions() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_update_actions",
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_UpdateAction.html`,
		Resolver:    fetchElasticacheUpdateAction,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticache"),
		Transform:   transformers.TransformWithStruct(&types.UpdateAction{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}
