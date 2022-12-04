// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"

func Armoperationalinsights() []Table {
	tables := []Table{
		{
      Name: "usage_metric",
      Struct: &armoperationalinsights.UsageMetric{},
      ResponseStruct: &armoperationalinsights.UsagesClientListResponse{},
      Client: &armoperationalinsights.UsagesClient{},
      ListFunc: (&armoperationalinsights.UsagesClient{}).NewListPager,
			NewFunc: armoperationalinsights.NewUsagesClient,
		},
		{
      Name: "cluster",
      Struct: &armoperationalinsights.Cluster{},
      ResponseStruct: &armoperationalinsights.ClustersClientListResponse{},
      Client: &armoperationalinsights.ClustersClient{},
      ListFunc: (&armoperationalinsights.ClustersClient{}).NewListPager,
			NewFunc: armoperationalinsights.NewClustersClient,
		},
		{
      Name: "workspace",
      Struct: &armoperationalinsights.Workspace{},
      ResponseStruct: &armoperationalinsights.DeletedWorkspacesClientListResponse{},
      Client: &armoperationalinsights.DeletedWorkspacesClient{},
      ListFunc: (&armoperationalinsights.DeletedWorkspacesClient{}).NewListPager,
			NewFunc: armoperationalinsights.NewDeletedWorkspacesClient,
		},
		{
      Name: "operation",
      Struct: &armoperationalinsights.Operation{},
      ResponseStruct: &armoperationalinsights.OperationsClientListResponse{},
      Client: &armoperationalinsights.OperationsClient{},
      ListFunc: (&armoperationalinsights.OperationsClient{}).NewListPager,
			NewFunc: armoperationalinsights.NewOperationsClient,
		},
		{
      Name: "management_group",
      Struct: &armoperationalinsights.ManagementGroup{},
      ResponseStruct: &armoperationalinsights.ManagementGroupsClientListResponse{},
      Client: &armoperationalinsights.ManagementGroupsClient{},
      ListFunc: (&armoperationalinsights.ManagementGroupsClient{}).NewListPager,
			NewFunc: armoperationalinsights.NewManagementGroupsClient,
		},
		{
      Name: "workspace",
      Struct: &armoperationalinsights.Workspace{},
      ResponseStruct: &armoperationalinsights.WorkspacesClientListResponse{},
      Client: &armoperationalinsights.WorkspacesClient{},
      ListFunc: (&armoperationalinsights.WorkspacesClient{}).NewListPager,
			NewFunc: armoperationalinsights.NewWorkspacesClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armoperationalinsights"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armoperationalinsights()...)
}