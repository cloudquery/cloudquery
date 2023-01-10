package mysql

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func server_configurations() *schema.Table {
	return &schema.Table{
		Name:      "azure_mysql_server_configurations",
		Resolver:  fetchServerConfigurations,
		Transform: transformers.TransformWithStruct(&armmysql.Configuration{}),
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
