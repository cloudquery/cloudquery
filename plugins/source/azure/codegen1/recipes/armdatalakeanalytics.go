// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics"

func Armdatalakeanalytics() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdatalakeanalytics.NewAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics",
		},
		{
			NewFunc: armdatalakeanalytics.NewComputePoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics",
		},
		{
			NewFunc: armdatalakeanalytics.NewDataLakeStoreAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics",
		},
		{
			NewFunc: armdatalakeanalytics.NewFirewallRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics",
		},
		{
			NewFunc: armdatalakeanalytics.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics",
		},
		{
			NewFunc: armdatalakeanalytics.NewLocationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics",
		},
		{
			NewFunc: armdatalakeanalytics.NewStorageAccountsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armdatalakeanalytics())
}