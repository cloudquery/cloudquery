package cosmos

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func sql_databases() *schema.Table {
	return &schema.Table{
		Name:      "azure_cosmos_sql_databases",
		Resolver:  fetchSqlDatabases,
		Transform: transformers.TransformWithStruct(&armcosmos.SQLDatabaseGetResults{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
