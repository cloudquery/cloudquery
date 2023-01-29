package postgresql

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func serverConfigurations() *schema.Table {
	return &schema.Table{
		Name:        "azure_postgresql_server_configurations",
		Resolver:    fetchServerConfigurations,
		Description: "https://learn.microsoft.com/en-us/rest/api/postgresql/singleserver/configurations/list-by-server?tabs=HTTP#configuration",
		Transform:   transformers.TransformWithStruct(&armpostgresql.Configuration{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{},
	}
}
