package mysql

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func server_configurations() *schema.Table {
	return &schema.Table{
		Name:                 "azure_mysql_server_configurations",
		Resolver:             fetchServerConfigurations,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/mysql/singleserver/configurations/list-by-server?tabs=HTTP#configuration",
		Transform:            transformers.TransformWithStruct(&armmysql.Configuration{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{},
	}
}
