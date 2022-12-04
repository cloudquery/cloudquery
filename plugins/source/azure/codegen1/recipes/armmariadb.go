// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"

func Armmariadb() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmariadb.NewLocationBasedRecommendedActionSessionsResultClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewPrivateEndpointConnectionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewLocationBasedRecommendedActionSessionsOperationStatusClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewLogFilesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewServerParametersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewDatabasesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewPrivateLinkResourcesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewAdvisorsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewServersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewFirewallRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewLocationBasedPerformanceTierClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewServerBasedPerformanceTierClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewRecommendedActionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewRecoverableServersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewTopQueryStatisticsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewWaitStatisticsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewVirtualNetworkRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewCheckNameAvailabilityClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewManagementClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewQueryTextsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewReplicasClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
		{
			NewFunc: armmariadb.NewServerSecurityAlertPoliciesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armmariadb())
}