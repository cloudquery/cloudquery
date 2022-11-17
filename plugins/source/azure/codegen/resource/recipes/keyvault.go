package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/keyvault"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func KeyVault() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct: new(armkeyvault.Vault),
			Resolver: &resource.FuncParams{
				Func: keyvault.VaultsClient.NewListBySubscriptionPager,
			},
			Children: []*resource.Resource{
				{
					Struct: new(armkeyvault.Key),
					Resolver: &resource.FuncParams{
						Func:   keyvault.KeysClient.NewListPager,
						Params: []string{"id.ResourceGroupName", "*vault.Name"},
					},
				},
				{
					Struct: new(armkeyvault.Secret),
					Resolver: &resource.FuncParams{
						Func:   keyvault.SecretsClient.NewListPager,
						Params: []string{"id.ResourceGroupName", "*vault.Name"},
					},
				},
			},
		},
		{
			Struct: new(armkeyvault.ManagedHsm),
			Resolver: &resource.FuncParams{
				Func: keyvault.ManagedHsmsClient.NewListBySubscriptionPager,
			},
		},
	}
}
