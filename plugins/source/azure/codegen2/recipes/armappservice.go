// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"

func Armappservice() []Table {
	tables := []Table{
		{
      Name: "certificate_order",
      Struct: &armappservice.CertificateOrder{},
      ResponseStruct: &armappservice.CertificateOrdersClientListResponse{},
      Client: &armappservice.CertificateOrdersClient{},
      ListFunc: (&armappservice.CertificateOrdersClient{}).NewListPager,
			NewFunc: armappservice.NewCertificateOrdersClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.CertificateRegistration/certificateOrders",
		},
		{
      Name: "app_certificate",
      Struct: &armappservice.AppCertificate{},
      ResponseStruct: &armappservice.CertificatesClientListResponse{},
      Client: &armappservice.CertificatesClient{},
      ListFunc: (&armappservice.CertificatesClient{}).NewListPager,
			NewFunc: armappservice.NewCertificatesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Web/certificates",
		},
		{
      Name: "deleted_site",
      Struct: &armappservice.DeletedSite{},
      ResponseStruct: &armappservice.DeletedWebAppsClientListResponse{},
      Client: &armappservice.DeletedWebAppsClient{},
      ListFunc: (&armappservice.DeletedWebAppsClient{}).NewListPager,
			NewFunc: armappservice.NewDeletedWebAppsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Web/deletedSites",
		},
		{
      Name: "domain",
      Struct: &armappservice.Domain{},
      ResponseStruct: &armappservice.DomainsClientListResponse{},
      Client: &armappservice.DomainsClient{},
      ListFunc: (&armappservice.DomainsClient{}).NewListPager,
			NewFunc: armappservice.NewDomainsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/domains",
		},
		{
      Name: "environment_resource",
      Struct: &armappservice.EnvironmentResource{},
      ResponseStruct: &armappservice.EnvironmentsClientListResponse{},
      Client: &armappservice.EnvironmentsClient{},
      ListFunc: (&armappservice.EnvironmentsClient{}).NewListPager,
			NewFunc: armappservice.NewEnvironmentsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Web/hostingEnvironments",
		},
		{
      Name: "plan",
      Struct: &armappservice.Plan{},
      ResponseStruct: &armappservice.PlansClientListResponse{},
      Client: &armappservice.PlansClient{},
      ListFunc: (&armappservice.PlansClient{}).NewListPager,
			NewFunc: armappservice.NewPlansClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Web/serverfarms",
		},
		{
      Name: "recommendation",
      Struct: &armappservice.Recommendation{},
      ResponseStruct: &armappservice.RecommendationsClientListResponse{},
      Client: &armappservice.RecommendationsClient{},
      ListFunc: (&armappservice.RecommendationsClient{}).NewListPager,
			NewFunc: armappservice.NewRecommendationsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Web/recommendations",
		},
		{
      Name: "resource_health_metadata",
      Struct: &armappservice.ResourceHealthMetadata{},
      ResponseStruct: &armappservice.ResourceHealthMetadataClientListResponse{},
      Client: &armappservice.ResourceHealthMetadataClient{},
      ListFunc: (&armappservice.ResourceHealthMetadataClient{}).NewListPager,
			NewFunc: armappservice.NewResourceHealthMetadataClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Web/resourceHealthMetadata",
		},
		{
      Name: "static_site_arm_resource",
      Struct: &armappservice.StaticSiteARMResource{},
      ResponseStruct: &armappservice.StaticSitesClientListResponse{},
      Client: &armappservice.StaticSitesClient{},
      ListFunc: (&armappservice.StaticSitesClient{}).NewListPager,
			NewFunc: armappservice.NewStaticSitesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Web/staticSites",
		},
		{
      Name: "top_level_domain",
      Struct: &armappservice.TopLevelDomain{},
      ResponseStruct: &armappservice.TopLevelDomainsClientListResponse{},
      Client: &armappservice.TopLevelDomainsClient{},
      ListFunc: (&armappservice.TopLevelDomainsClient{}).NewListPager,
			NewFunc: armappservice.NewTopLevelDomainsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains",
		},
		{
      Name: "site",
      Struct: &armappservice.Site{},
      ResponseStruct: &armappservice.WebAppsClientListResponse{},
      Client: &armappservice.WebAppsClient{},
      ListFunc: (&armappservice.WebAppsClient{}).NewListPager,
			NewFunc: armappservice.NewWebAppsClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Web/sites",
		},
	}

	for i := range tables {
		tables[i].Service = "armappservice"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armappservice()...)
}