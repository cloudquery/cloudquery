// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"

func Armresources() []Table {
	tables := []Table{
		{
      Name: "deployment_operation",
      Struct: &armresources.DeploymentOperation{},
      ResponseStruct: &armresources.DeploymentOperationsClientListResponse{},
      Client: &armresources.DeploymentOperationsClient{},
      ListFunc: (&armresources.DeploymentOperationsClient{}).NewListPager,
			NewFunc: armresources.NewDeploymentOperationsClient,
		},
		{
      Name: "operation",
      Struct: &armresources.Operation{},
      ResponseStruct: &armresources.OperationsClientListResponse{},
      Client: &armresources.OperationsClient{},
      ListFunc: (&armresources.OperationsClient{}).NewListPager,
			NewFunc: armresources.NewOperationsClient,
		},
		{
      Name: "generic_resource_expanded",
      Struct: &armresources.GenericResourceExpanded{},
      ResponseStruct: &armresources.ClientListResponse{},
      Client: &armresources.Client{},
      ListFunc: (&armresources.Client{}).NewListPager,
			NewFunc: armresources.NewClient,
		},
		{
      Name: "resource_group",
      Struct: &armresources.ResourceGroup{},
      ResponseStruct: &armresources.ResourceGroupsClientListResponse{},
      Client: &armresources.ResourceGroupsClient{},
      ListFunc: (&armresources.ResourceGroupsClient{}).NewListPager,
			NewFunc: armresources.NewResourceGroupsClient,
		},
		{
      Name: "provider",
      Struct: &armresources.Provider{},
      ResponseStruct: &armresources.ProvidersClientListResponse{},
      Client: &armresources.ProvidersClient{},
      ListFunc: (&armresources.ProvidersClient{}).NewListPager,
			NewFunc: armresources.NewProvidersClient,
		},
		{
      Name: "tag_details",
      Struct: &armresources.TagDetails{},
      ResponseStruct: &armresources.TagsClientListResponse{},
      Client: &armresources.TagsClient{},
      ListFunc: (&armresources.TagsClient{}).NewListPager,
			NewFunc: armresources.NewTagsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armresources"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armresources()...)
}