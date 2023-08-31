package sql

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func serverAdmins() *schema.Table {
	return &schema.Table{
		Name:                 "azure_sql_server_admins",
		Resolver:             fetchServerAdmins,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/sql/2020-08-01-preview/server-azure-ad-administrators/list-by-server?tabs=HTTP#serverazureadadministrator",
		Transform:            transformers.TransformWithStruct(&armsql.ServerAzureADAdministrator{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}
