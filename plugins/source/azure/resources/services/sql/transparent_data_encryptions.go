package sql

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func transparentDataEncryptions() *schema.Table {
	return &schema.Table{
		Name:        "azure_sql_transparent_data_encryptions",
		Resolver:    fetchTransparentDataEncryptions,
		Description: "https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/transparent-data-encryptions/list-by-database?tabs=HTTP#logicaldatabasetransparentdataencryption",
		Transform:   transformers.TransformWithStruct(&armsql.LogicalDatabaseTransparentDataEncryption{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}
