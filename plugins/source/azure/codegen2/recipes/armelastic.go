// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"

func Armelastic() []Table {
	tables := []Table{
		{
      Name: "vm_resources",
      Struct: &armelastic.VMResources{},
      ResponseStruct: &armelastic.VMHostClientListResponse{},
      Client: &armelastic.VMHostClient{},
      ListFunc: (&armelastic.VMHostClient{}).NewListPager,
			NewFunc: armelastic.NewVMHostClient,
		},
		{
      Name: "monitoring_tag_rules",
      Struct: &armelastic.MonitoringTagRules{},
      ResponseStruct: &armelastic.TagRulesClientListResponse{},
      Client: &armelastic.TagRulesClient{},
      ListFunc: (&armelastic.TagRulesClient{}).NewListPager,
			NewFunc: armelastic.NewTagRulesClient,
		},
		{
      Name: "monitored_resource",
      Struct: &armelastic.MonitoredResource{},
      ResponseStruct: &armelastic.MonitoredResourcesClientListResponse{},
      Client: &armelastic.MonitoredResourcesClient{},
      ListFunc: (&armelastic.MonitoredResourcesClient{}).NewListPager,
			NewFunc: armelastic.NewMonitoredResourcesClient,
		},
		{
      Name: "monitor_resource",
      Struct: &armelastic.MonitorResource{},
      ResponseStruct: &armelastic.MonitorsClientListResponse{},
      Client: &armelastic.MonitorsClient{},
      ListFunc: (&armelastic.MonitorsClient{}).NewListPager,
			NewFunc: armelastic.NewMonitorsClient,
		},
		{
      Name: "operation_result",
      Struct: &armelastic.OperationResult{},
      ResponseStruct: &armelastic.OperationsClientListResponse{},
      Client: &armelastic.OperationsClient{},
      ListFunc: (&armelastic.OperationsClient{}).NewListPager,
			NewFunc: armelastic.NewOperationsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armelastic"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armelastic()...)
}