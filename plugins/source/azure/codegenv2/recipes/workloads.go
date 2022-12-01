package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"

func WorkloadsResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "monitors",
			Struct: &armworkloads.Monitor{},
			ResponseStruct: &armworkloads.MonitorsClientListResponse{},
			Client: &armworkloads.MonitorsClient{},
			ListFunc: (&armworkloads.MonitorsClient{}).NewListPager,
			NewFunc: armworkloads.NewMonitorsClient,
			OutputField: "Value",
		},
		{
			SubService: "operations",
			Struct: &armworkloads.Operation{},
			ResponseStruct: &armworkloads.OperationsClientListResponse{},
			Client: &armworkloads.OperationsClient{},
			ListFunc: (&armworkloads.OperationsClient{}).NewListPager,
			NewFunc: armworkloads.NewOperationsClient,
			OutputField: "Value",
		},
		{
			SubService: "php_workloads",
			Struct: &armworkloads.PhpWorkloadResource{},
			ResponseStruct: &armworkloads.PhpWorkloadsClientListByResourceGroupResponse{},
			Client: &armworkloads.PhpWorkloadsClient{},
			ListFunc: (&armworkloads.PhpWorkloadsClient{}).NewListBySubscriptionPager,
			NewFunc: armworkloads.NewPhpWorkloadsClient,
			OutputField: "Value",
		},
		{
			SubService: "provider_instances",
			Struct: &armworkloads.ProviderInstance{},
			ResponseStruct: &armworkloads.ProviderInstancesClientListResponse{},
			Client: &armworkloads.ProviderInstancesClient{},
			ListFunc: (&armworkloads.OperationsClient{}).NewListPager,
			NewFunc: armworkloads.NewProviderInstancesClient,
			OutputField: "Value",
		},
		{
			SubService: "sap_application_server_instances",
			Struct: &armworkloads.SAPApplicationServerInstance{},
			ResponseStruct: &armworkloads.SAPApplicationServerInstancesClientListResponse{},
			Client: &armworkloads.SAPApplicationServerInstancesClient{},
			ListFunc: (&armworkloads.SAPApplicationServerInstancesClient{}).NewListPager,
			NewFunc: armworkloads.NewSAPApplicationServerInstancesClient,
			OutputField: "Value",
		},
		{
			SubService: "sap_central_server_instances",
			Struct: &armworkloads.SAPCentralServerInstance{},
			ResponseStruct: &armworkloads.SAPCentralInstancesClientListResponse{},
			Client: &armworkloads.SAPCentralInstancesClient{},
			ListFunc: (&armworkloads.SAPCentralInstancesClient{}).NewListPager,
			NewFunc: armworkloads.NewSAPCentralInstancesClient,
			OutputField: "Value",
		},
		{
			SubService: "sap_database_instances",
			Struct: &armworkloads.SAPDatabaseInstance{},
			ResponseStruct: &armworkloads.SAPDatabaseInstancesClientListResponse{},
			Client: &armworkloads.SAPDatabaseInstancesClient{},
			ListFunc: (&armworkloads.SAPDatabaseInstancesClient{}).NewListPager,
			NewFunc: armworkloads.NewSAPDatabaseInstancesClient,
			OutputField: "Value",
		},
		{
			SubService: "sap_virtual_instances",
			Struct: &armworkloads.SAPVirtualInstance{},
			ResponseStruct: &armworkloads.SAPVirtualInstancesClientListBySubscriptionResponse{},
			Client: &armworkloads.SAPVirtualInstancesClient{},
			ListFunc: (&armworkloads.SAPVirtualInstancesClient{}).NewListBySubscriptionPager,
			NewFunc: armworkloads.NewSAPVirtualInstancesClient,
			OutputField: "Value",
		},
		{
			SubService: "sap_skus",
			Struct: &armworkloads.SKU{},
			ResponseStruct: &armworkloads.SKUsClientListResponse{},
			Client: &armworkloads.SKUsClient{},
			ListFunc: (&armworkloads.SKUsClient{}).NewListPager,
			NewFunc: armworkloads.NewSKUsClient,
			OutputField: "Value",
		},
	}

	for _, r := range resources {
		r.ImportPath = "workloads/armworkloads"
		r.Service = "armworkloads"
		r.Template = "list"
	}

	return resources
}