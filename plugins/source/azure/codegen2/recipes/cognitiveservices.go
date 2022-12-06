// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"

func Armcognitiveservices() []Table {
	tables := []Table{
		{
			Name:           "account",
			Struct:         &armcognitiveservices.Account{},
			ResponseStruct: &armcognitiveservices.AccountsClientListResponse{},
			Client:         &armcognitiveservices.AccountsClient{},
			ListFunc:       (&armcognitiveservices.AccountsClient{}).NewListPager,
			NewFunc:        armcognitiveservices.NewAccountsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/accounts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.CognitiveServices")`,
		},
		{
			Name:           "account",
			Struct:         &armcognitiveservices.Account{},
			ResponseStruct: &armcognitiveservices.DeletedAccountsClientListResponse{},
			Client:         &armcognitiveservices.DeletedAccountsClient{},
			ListFunc:       (&armcognitiveservices.DeletedAccountsClient{}).NewListPager,
			NewFunc:        armcognitiveservices.NewDeletedAccountsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/deletedAccounts",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.CognitiveServices")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armcognitiveservices"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armcognitiveservices()...)
}
