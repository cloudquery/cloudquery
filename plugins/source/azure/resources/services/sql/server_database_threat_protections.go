package sql

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func serverDatabaseThreatProtections() *schema.Table {
	return &schema.Table{
		Name:        "azure_sql_server_database_threat_protections",
		Resolver:    fetchDatabaseThreatProtections,
		Description: "https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/database-advanced-threat-protection-settings/list-by-database?tabs=HTTP#databaseadvancedthreatprotection",
		Transform:   transformers.TransformWithStruct(&armsql.DatabaseAdvancedThreatProtection{}),
		Columns:     schema.ColumnList{client.SubscriptionID, client.IDColumn},
	}
}
