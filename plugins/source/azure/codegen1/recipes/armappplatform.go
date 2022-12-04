// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"

func Armappplatform() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armappplatform.NewBindingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
		{
			NewFunc: armappplatform.NewBuildServiceClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
		{
			NewFunc: armappplatform.NewMonitoringSettingsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
		{
			NewFunc: armappplatform.NewServiceRegistriesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
		{
			NewFunc: armappplatform.NewServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
		{
			NewFunc: armappplatform.NewSKUsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
		{
			NewFunc: armappplatform.NewAppsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
		{
			NewFunc: armappplatform.NewCertificatesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
		{
			NewFunc: armappplatform.NewConfigServersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
		{
			NewFunc: armappplatform.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
		{
			NewFunc: armappplatform.NewRuntimeVersionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
		{
			NewFunc: armappplatform.NewBuildpackBindingClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
		{
			NewFunc: armappplatform.NewBuildServiceAgentPoolClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
		{
			NewFunc: armappplatform.NewCustomDomainsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
		{
			NewFunc: armappplatform.NewDeploymentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
		{
			NewFunc: armappplatform.NewBuildServiceBuilderClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
		{
			NewFunc: armappplatform.NewConfigurationServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armappplatform())
}