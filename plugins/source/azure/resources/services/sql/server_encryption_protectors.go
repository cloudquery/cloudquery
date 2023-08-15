package sql

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func serverEncryptionProtectors() *schema.Table {
	return &schema.Table{
		Name:                 "azure_sql_server_encryption_protectors",
		Resolver:             fetchEncryptionProtectors,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/sql/2021-11-01/encryption-protectors/list-by-server?tabs=HTTP#encryptionprotector",
		Transform:            transformers.TransformWithStruct(&armsql.EncryptionProtector{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}
