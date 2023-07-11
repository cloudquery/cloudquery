package sql

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func serverDatabaseBlobAuditingPolicies() *schema.Table {
	return &schema.Table{
		Name:                 "azure_sql_server_database_blob_auditing_policies",
		Resolver:             fetchDatabaseBlobAuditingPolicies,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/database-blob-auditing-policies/list-by-database?tabs=HTTP#databaseblobauditingpolicy",
		Transform:            transformers.TransformWithStruct(&armsql.DatabaseBlobAuditingPolicy{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}
