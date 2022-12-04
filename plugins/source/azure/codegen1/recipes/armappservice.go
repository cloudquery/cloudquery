// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"

func Armappservice() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armappservice.NewCertificateRegistrationProviderClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewCertificatesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewRecommendationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewContainerAppsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewDeletedWebAppsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewDiagnosticsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewGlobalClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewWebAppsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewEnvironmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewDomainsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewProviderClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewCertificateOrdersDiagnosticsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewResourceHealthMetadataClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewContainerAppsRevisionsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewStaticSitesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewTopLevelDomainsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewCertificateOrdersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewDomainRegistrationProviderClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewKubeEnvironmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewPlansClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
		{
			NewFunc: armappservice.NewWebSiteManagementClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armappservice())
}