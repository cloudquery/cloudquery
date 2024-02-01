package glue

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func databaseTableIndexes() *schema.Table {
	tableName := "aws_glue_database_table_indexes"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_PartitionIndexDescriptor.html`,
		Resolver:    fetchGlueDatabaseTableIndexes,
		Transform:   transformers.TransformWithStruct(&types.PartitionIndexDescriptor{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "database_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("database_arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "database_table_name",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("name"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "index_name",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("IndexName"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchGlueDatabaseTableIndexes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGlue).Glue
	d := parent.Parent.Item.(types.Database)
	t := parent.Item.(types.Table)
	input := glue.GetPartitionIndexesInput{DatabaseName: d.Name, CatalogId: d.CatalogId, TableName: t.Name}
	paginator := glue.NewGetPartitionIndexesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *glue.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.PartitionIndexDescriptorList
	}
	return nil
}
