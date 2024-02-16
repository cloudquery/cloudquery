package dynamodb

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
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
		Transform:   transformers.TransformWithStruct(&types.ReplicaAutoScalingDescription{}, transformers.WithPrimaryKeyComponents("RegionName")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "table_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchDynamodbTableReplicaAutoScalings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	par := parent.Item.(*types.TableDescription)

	if aws.ToString(par.GlobalTableVersion) != "2019.11.21" {
		// "This operation only applies to Version 2019.11.21 of global tables"
		return nil
	}

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDynamodb).Dynamodb

	output, err := svc.DescribeTableReplicaAutoScaling(ctx, &dynamodb.DescribeTableReplicaAutoScalingInput{
		TableName: par.TableName,
	}, func(options *dynamodb.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	res <- output.TableAutoScalingDescription.Replicas
	return nil
}
