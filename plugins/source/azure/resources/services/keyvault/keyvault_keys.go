package keyvault

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func keyvault_keys() *schema.Table {
	return &schema.Table{
		Name:        "azure_keyvault_keyvault_keys",
		Resolver:    fetchKeyvaultKeys,
		Description: "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault@v1.0.0#Key",
		Transform:   transformers.TransformWithStruct(&armkeyvault.Key{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{},
	}
}
