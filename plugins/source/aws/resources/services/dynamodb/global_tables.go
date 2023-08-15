package dynamodb

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func GlobalTables() *schema.Table {
	tableName := "aws_dynamodb_global_tables"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GlobalTableDescription.html
This table only contains version 2017.11.29 (Legacy) Global Tables. See aws_dynamodb_tables for version 2019.11.21 (Current) Global Tables.
`,
		Resolver:            fetchGlobalTables,
		PreResourceResolver: getGlobalTable,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "dynamodb"),
		Transform:           transformers.TransformWithStruct(&types.GlobalTableDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(true),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("GlobalTableArn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveDynamodbGlobalTableTags,
			},
		},
	}
}

func fetchGlobalTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Dynamodb

	config := dynamodb.ListGlobalTablesInput{
		RegionName: aws.String(cl.Region),
	}
	// No paginator available
	for {
		output, err := svc.ListGlobalTables(ctx, &config, func(options *dynamodb.Options) {
			options.Region = cl.Region
		})
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
	cl := meta.(*client.Client)
	svc := cl.Services().Dynamodb

	table := resource.Item.(types.GlobalTable)

	response, err := svc.DescribeGlobalTable(ctx, &dynamodb.DescribeGlobalTableInput{GlobalTableName: table.GlobalTableName}, func(options *dynamodb.Options) {
		options.Region = cl.Region
	})
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
	var tags []types.Tag
	input := &dynamodb.ListTagsOfResourceInput{
		ResourceArn: table.GlobalTableArn,
	}
	// // No paginator available
	for {
		response, err := svc.ListTagsOfResource(ctx, input, func(options *dynamodb.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		tags = append(tags, response.Tags...)
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, client.TagsToMap(tags))
}
