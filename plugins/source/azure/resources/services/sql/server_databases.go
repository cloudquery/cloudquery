package sql

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func serverDatabases() *schema.Table {
	return &schema.Table{
		Name:                 "azure_sql_server_databases",
		Resolver:             fetchDatabases,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/databases/list-by-server?tabs=HTTP#database",
		Transform:            transformers.TransformWithStruct(&armsql.Database{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},

		Relations: []*schema.Table{
			serverDatabaseBlobAuditingPolicies(),
			transparentDataEncryptions(),
			serverDatabaseThreatProtections(),
			databaseVulnerabilityAssessments(),
			longTermRetentionPolicies(),
		},
	}
}
