package athena

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func dataCatalogDatabases() *schema.Table {
	tableName := "aws_athena_data_catalog_databases"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/athena/latest/APIReference/API_Database.html`,
		Resolver:    fetchAthenaDataCatalogDatabases,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "athena"),
		Transform:   transformers.TransformWithStruct(&types.Database{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "data_catalog_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
			{
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Name"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			dataCatalogDatabaseTables(),
		},
	}
}

func fetchAthenaDataCatalogDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena
	input := athena.ListDatabasesInput{
		CatalogName: parent.Item.(types.DataCatalog).Name,
	}
	paginator := athena.NewListDatabasesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *athena.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.DatabaseList
	}
	return nil
}
