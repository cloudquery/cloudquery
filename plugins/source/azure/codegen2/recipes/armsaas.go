// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/saas/armsaas"

func Armsaas() []Table {
	tables := []Table{
		{
      Name: "resource",
      Struct: &armsaas.Resource{},
      ResponseStruct: &armsaas.ResourcesClientListResponse{},
      Client: &armsaas.ResourcesClient{},
      ListFunc: (&armsaas.ResourcesClient{}).NewListPager,
			NewFunc: armsaas.NewResourcesClient,
		},
		{
      Name: "app",
      Struct: &armsaas.App{},
      ResponseStruct: &armsaas.ApplicationsClientListResponse{},
      Client: &armsaas.ApplicationsClient{},
      ListFunc: (&armsaas.ApplicationsClient{}).NewListPager,
			NewFunc: armsaas.NewApplicationsClient,
		},
		{
      Name: "app_operation",
      Struct: &armsaas.AppOperation{},
      ResponseStruct: &armsaas.OperationsClientListResponse{},
      Client: &armsaas.OperationsClient{},
      ListFunc: (&armsaas.OperationsClient{}).NewListPager,
			NewFunc: armsaas.NewOperationsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armsaas"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armsaas()...)
}