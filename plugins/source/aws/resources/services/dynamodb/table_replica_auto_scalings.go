package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TableReplicaAutoScalings() *schema.Table {
	return &schema.Table{
		Name:        "aws_dynamodb_table_replica_auto_scalings",
		Description: `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ReplicaAutoScalingDescription.html`,
		Resolver:    fetchDynamodbTableReplicaAutoScalings,
		Multiplex:   client.ServiceAccountRegionMultiplexer("dynamodb"),
		Transform:   transformers.TransformWithStruct(&types.ReplicaAutoScalingDescription{}),
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
				Name:     "table_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
