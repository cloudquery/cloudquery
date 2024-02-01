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

func databaseTables() *schema.Table {
	tableName := "aws_glue_database_tables"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_Table.html`,
		Resolver:    fetchGlueDatabaseTables,
		Transform:   transformers.TransformWithStruct(&types.Table{}, transformers.WithPrimaryKeyComponents("Name")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "database_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},

		Relations: []*schema.Table{
			databaseTableIndexes(),
		},
	}
}

func fetchGlueDatabaseTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Database)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGlue).Glue
	input := glue.GetTablesInput{
		DatabaseName: r.Name,
	}
	paginator := glue.NewGetTablesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *glue.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.TableList
	}
	return nil
}
