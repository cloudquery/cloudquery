package glue

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Databases() *schema.Table {
	tableName := "aws_glue_databases"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_Database.html`,
		Resolver:    fetchGlueDatabases,
		Transform:   transformers.TransformWithStruct(&types.Database{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveGlueDatabaseArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueDatabaseTags,
			},
		},

		Relations: []*schema.Table{
			databaseTables(),
		},
	}
}

func fetchGlueDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetDatabasesInput{}
	for {
		result, err := svc.GetDatabases(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.DatabaseList
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
func resolveGlueDatabaseArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, databaseARN(cl, aws.ToString(resource.Item.(types.Database).Name)))
}
func resolveGlueDatabaseTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetTagsInput{
		ResourceArn: aws.String(databaseARN(cl, aws.ToString(resource.Item.(types.Database).Name))),
	}

	response, err := svc.GetTags(ctx, &input)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, response.Tags)
}
func fetchGlueDatabaseTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Database)
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetTablesInput{
		DatabaseName: r.Name,
	}
	for {
		result, err := svc.GetTables(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.TableList
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}
func fetchGlueDatabaseTableIndexes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	d := parent.Parent.Item.(types.Database)
	t := parent.Item.(types.Table)
	input := glue.GetPartitionIndexesInput{DatabaseName: d.Name, CatalogId: d.CatalogId, TableName: t.Name}
	for {
		result, err := svc.GetPartitionIndexes(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.PartitionIndexDescriptorList
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}

func databaseARN(cl *client.Client, name string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.GlueService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("database/%s", name),
	}.String()
}
