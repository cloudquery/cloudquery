package athena

import (
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DataCatalogDatabases() *schema.Table {
	tableName := "aws_athena_data_catalog_databases"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/athena/latest/APIReference/API_Database.html`,
		Resolver:    fetchAthenaDataCatalogDatabases,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "athena"),
		Transform:   client.TransformWithStruct(&types.Database{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "data_catalog_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			DataCatalogDatabaseTables(),
		},
	}
}
