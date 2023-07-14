package keyvault

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func keyvault_keys() *schema.Table {
	return &schema.Table{
		Name:                 "azure_keyvault_keyvault_keys",
		Resolver:             fetchKeyvaultKeys,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault@v1.0.0#Key",
		Transform:            transformers.TransformWithStruct(&armkeyvault.Key{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{},
	}
}
