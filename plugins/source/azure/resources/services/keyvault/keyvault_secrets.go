package keyvault

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func keyvault_secrets() *schema.Table {
	return &schema.Table{
		Name:        "azure_keyvault_keyvault_secrets",
		Resolver:    fetchKeyvaultSecrets,
		Description: "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault@v1.0.0#Secret",
		Transform:   transformers.TransformWithStruct(&armkeyvault.Secret{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{},
	}
}
