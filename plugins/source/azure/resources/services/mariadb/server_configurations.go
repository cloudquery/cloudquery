package mariadb

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func server_configurations() *schema.Table {
	return &schema.Table{
		Name:        "azure_mariadb_server_configurations",
		Resolver:    fetchServerConfigurations,
		Description: "https://learn.microsoft.com/en-us/rest/api/mariadb/configurations/list-by-server?tabs=HTTP#configuration",
		Transform:   transformers.TransformWithStruct(&armmariadb.Configuration{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{},
	}
}
