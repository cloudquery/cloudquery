package dynamodb

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func tableReplicaAutoScalings() *schema.Table {
	tableName := "aws_dynamodb_table_replica_auto_scalings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ReplicaAutoScalingDescription.html`,
		Resolver:    fetchDynamodbTableReplicaAutoScalings,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "dynamodb"),
		Transform:   transformers.TransformWithStruct(&types.ReplicaAutoScalingDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "table_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
