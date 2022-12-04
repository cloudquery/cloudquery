// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"

func Armpostgresqlhsc() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armpostgresqlhsc.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc",
		},
		{
			NewFunc: armpostgresqlhsc.NewRolesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc",
		},
		{
			NewFunc: armpostgresqlhsc.NewServersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc",
		},
		{
			NewFunc: armpostgresqlhsc.NewConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc",
		},
		{
			NewFunc: armpostgresqlhsc.NewFirewallRulesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc",
		},
		{
			NewFunc: armpostgresqlhsc.NewServerGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armpostgresqlhsc())
}