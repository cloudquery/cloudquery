package cosmos

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func sql_databases() *schema.Table {
	return &schema.Table{
		Name:                 "azure_cosmos_sql_databases",
		Resolver:             fetchSqlDatabases,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/cosmos-db-resource-provider/2022-05-15/sql-resources/list-sql-databases?tabs=HTTP#sqldatabasegetresults",
		Transform:            transformers.TransformWithStruct(&armcosmos.SQLDatabaseGetResults{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{},
	}
}
