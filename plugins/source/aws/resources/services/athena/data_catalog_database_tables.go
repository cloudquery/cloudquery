package athena

import (
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DataCatalogDatabaseTables() *schema.Table {
	return &schema.Table{
		Name:        "aws_athena_data_catalog_database_tables",
		Description: `https://docs.aws.amazon.com/athena/latest/APIReference/API_TableMetadata.html`,
		Resolver:    fetchAthenaDataCatalogDatabaseTables,
		Multiplex:   client.ServiceAccountRegionMultiplexer("athena"),
		Transform:   transformers.TransformWithStruct(&types.TableMetadata{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
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
