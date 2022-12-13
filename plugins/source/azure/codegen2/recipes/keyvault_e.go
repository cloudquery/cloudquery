package recipes

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

func init() {
	tables := []Table{
		{
			Service:        "armkeyvault",
			Name:           "keyvault_managed_hsms",
			Struct:         &armkeyvault.ManagedHsm{},
			ResponseStruct: &armkeyvault.ManagedHsmsClientListBySubscriptionResponse{},
			Client:         &armkeyvault.ManagedHsmsClient{},
			ListFunc:       (&armkeyvault.ManagedHsmsClient{}).NewListBySubscriptionPager,
			NewFunc:        armkeyvault.NewManagedHsmsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/managedHSMs",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_KeyVault)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armkeyvault",
			Name:           "keyvault",
			Struct:         &armkeyvault.Resource{},
			ResponseStruct: &armkeyvault.VaultsClientListResponse{},
			Client:         &armkeyvault.VaultsClient{},
			ListFunc:       (&armkeyvault.VaultsClient{}).NewListPager,
			NewFunc:        armkeyvault.NewVaultsClient,
			URL:            "/subscriptions/{subscriptionId}/resources",
			Multiplex:      `client.SubscriptionMultiplex`,
			ExtraColumns:   DefaultExtraColumns,
			Relations: []*Table{
				{
					Service:        "armkeyvault",
					Name:           "keyvault_keys",
					Struct:         &armkeyvault.Key{},
					ResponseStruct: &armkeyvault.KeysClientListResponse{},
					Client:         &armkeyvault.KeysClient{},
					ListFunc:       (&armkeyvault.KeysClient{}).NewListPager,
					NewFunc:        armkeyvault.NewKeysClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/keys",
					SkipFetch:      true,
				},
				{
					Service:        "armkeyvault",
					Name:           "keyvault_secrets",
					Struct:         &armkeyvault.Secret{},
					ResponseStruct: &armkeyvault.SecretsClientListResponse{},
					Client:         &armkeyvault.SecretsClient{},
					ListFunc:       (&armkeyvault.SecretsClient{}).NewListPager,
					NewFunc:        armkeyvault.NewSecretsClient,
					URL:            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/secrets",
					SkipFetch:      true,
				},
			},
		},
	}
	Tables = append(Tables, tables...)
}
