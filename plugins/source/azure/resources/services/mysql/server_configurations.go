package mysql

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func server_configurations() *schema.Table {
	return &schema.Table{
		Name:        "azure_mysql_server_configurations",
		Resolver:    fetchServerConfigurations,
		Description: "https://learn.microsoft.com/en-us/rest/api/mysql/singleserver/configurations/list-by-server?tabs=HTTP#configuration",
		Transform:   transformers.TransformWithStruct(&armmysql.Configuration{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{},
	}
}
