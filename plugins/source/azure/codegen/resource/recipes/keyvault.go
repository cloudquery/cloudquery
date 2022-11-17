package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/keyvault"
	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/resource"
)

func KeyVault() []*resource.Resource {
	return []*resource.Resource{
		{
			Struct:   new(armkeyvault.Vault),
			Resolver: keyvault.VaultsClient.NewListBySubscriptionPager,
			Children: []*resource.Resource{
				{
					Struct:   new(armkeyvault.Key),
					Resolver: keyvault.KeysClient.NewListPager,
				},
				{
					Struct:   new(armkeyvault.Secret),
					Resolver: keyvault.SecretsClient.NewListPager,
				},
			},
		},
		{
			Struct:   new(armkeyvault.ManagedHsm),
			Resolver: keyvault.ManagedHsmsClient.NewListBySubscriptionPager,
		},
	}
}
