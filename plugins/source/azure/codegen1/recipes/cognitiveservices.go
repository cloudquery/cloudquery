// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"

func Armcognitiveservices() []*Table {
	tables := []*Table{
		{
			NewFunc: armcognitiveservices.NewDeletedAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/deletedAccounts",
		},
		{
			NewFunc: armcognitiveservices.NewResourceSKUsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/skus",
		},
		{
			NewFunc: armcognitiveservices.NewAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/accounts",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armcognitiveservices())
}