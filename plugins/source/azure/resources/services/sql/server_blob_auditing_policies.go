package sql

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func server_blob_auditing_policies() *schema.Table {
	return &schema.Table{
		Name:                 "azure_sql_server_blob_auditing_policies",
		Resolver:             fetchServerBlobAuditingPolicies,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/server-blob-auditing-policies/list-by-server?tabs=HTTP#serverblobauditingpolicy",
		Transform:            transformers.TransformWithStruct(&armsql.ServerBlobAuditingPolicy{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}
