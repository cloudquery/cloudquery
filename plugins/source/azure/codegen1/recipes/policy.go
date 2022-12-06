// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"

func Armpolicy() []*Table {
	tables := []*Table{
		{
			NewFunc: armpolicy.NewDefinitionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions",
		},
		{
			NewFunc: armpolicy.NewExemptionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyExemptions",
		},
		{
			NewFunc: armpolicy.NewAssignmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyAssignments",
		},
		{
			NewFunc: armpolicy.NewDataPolicyManifestsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy",
			URL: "/providers/Microsoft.Authorization/dataPolicyManifests",
		},
		{
			NewFunc: armpolicy.NewSetDefinitionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policySetDefinitions",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armpolicy())
}