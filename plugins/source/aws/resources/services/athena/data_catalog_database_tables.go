package athena

import (
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DataCatalogDatabaseTables() *schema.Table {
	tableName := "aws_athena_data_catalog_database_tables"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/athena/latest/APIReference/API_TableMetadata.html`,
		Resolver:    fetchAthenaDataCatalogDatabaseTables,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "athena"),
		Transform:   client.TransformWithStruct(&types.TableMetadata{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "data_catalog_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("data_catalog_arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "data_catalog_database_name",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("name"),
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
	}
}
