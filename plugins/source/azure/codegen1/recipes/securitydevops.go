// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops"

func Armsecuritydevops() []*Table {
	tables := []*Table{
		{
			NewFunc: armsecuritydevops.NewAzureDevOpsOrgClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/azureDevOpsConnectors/{azureDevOpsConnectorName}/orgs",
		},
		{
			NewFunc: armsecuritydevops.NewGitHubOwnerClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securitydevops/armsecuritydevops",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SecurityDevOps/gitHubConnectors/{gitHubConnectorName}/owners",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armsecuritydevops())
}