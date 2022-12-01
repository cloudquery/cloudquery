package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

func KeyVaultResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "vaults",
			Struct: &armkeyvault.Vault{},
			ResponseStruct: &armkeyvault.VaultsClientListBySubscriptionResponse{},
			Client: &armkeyvault.VaultsClient{},
			ListFunc: (&armkeyvault.VaultsClient{}).NewListBySubscriptionPager,
			NewFunc: armkeyvault.NewVaultsClient,
			OutputField: "Value",
		},
	}

	for _, r := range resources {
		r.ImportPath = "keyvault/armkeyvault"
		r.Service = "armkeyvault"
		r.Template = "list"
	}

	return resources
}