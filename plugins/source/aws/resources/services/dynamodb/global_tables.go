package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func GlobalTables() *schema.Table {
	tableName := "aws_dynamodb_global_tables"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GlobalTableDescription.html`,
		Resolver:            fetchGlobalTables,
		PreResourceResolver: getGlobalTable,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "dynamodb"),
		Transform:           transformers.TransformWithStruct(&types.TableDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GlobalTableArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveDynamodbGlobalTableTags,
			},
		},
		Relations: []*schema.Table{},
	}
}

func fetchGlobalTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Dynamodb

	config := dynamodb.ListGlobalTablesInput{
		RegionName: aws.String(c.Region),
	}
	for {
		output, err := svc.ListGlobalTables(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.GlobalTables

		if aws.ToString(output.LastEvaluatedGlobalTableName) == "" {
			break
		}
		config.ExclusiveStartGlobalTableName = output.LastEvaluatedGlobalTableName
	}

	return nil
}

func getGlobalTable(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client).Services().Dynamodb

	table := resource.Item.(types.GlobalTable)

	response, err := svc.DescribeGlobalTable(ctx, &dynamodb.DescribeGlobalTableInput{GlobalTableName: table.GlobalTableName})
	if err != nil {
		return err
	}

	resource.Item = response.GlobalTableDescription
	return nil
}

func resolveDynamodbGlobalTableTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	table := resource.Item.(*types.GlobalTableDescription)

	cl := meta.(*client.Client)
	svc := cl.Services().Dynamodb
	response, err := svc.ListTagsOfResource(ctx, &dynamodb.ListTagsOfResourceInput{
		ResourceArn: table.GlobalTableArn,
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(response.Tags))
}
