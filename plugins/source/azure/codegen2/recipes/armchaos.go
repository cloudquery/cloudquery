// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"

func Armchaos() []Table {
	tables := []Table{
		{
      Name: "capability",
      Struct: &armchaos.Capability{},
      ResponseStruct: &armchaos.CapabilitiesClientListResponse{},
      Client: &armchaos.CapabilitiesClient{},
      ListFunc: (&armchaos.CapabilitiesClient{}).NewListPager,
			NewFunc: armchaos.NewCapabilitiesClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{parentProviderNamespace}/{parentResourceType}/{parentResourceName}/providers/Microsoft.Chaos/targets/{targetName}/capabilities",
		},
		{
      Name: "capability_type",
      Struct: &armchaos.CapabilityType{},
      ResponseStruct: &armchaos.CapabilityTypesClientListResponse{},
      Client: &armchaos.CapabilityTypesClient{},
      ListFunc: (&armchaos.CapabilityTypesClient{}).NewListPager,
			NewFunc: armchaos.NewCapabilityTypesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Chaos/locations/{locationName}/targetTypes/{targetTypeName}/capabilityTypes",
		},
		{
      Name: "experiment",
      Struct: &armchaos.Experiment{},
      ResponseStruct: &armchaos.ExperimentsClientListResponse{},
      Client: &armchaos.ExperimentsClient{},
      ListFunc: (&armchaos.ExperimentsClient{}).NewListPager,
			NewFunc: armchaos.NewExperimentsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Chaos/experiments",
		},
		{
      Name: "target_type",
      Struct: &armchaos.TargetType{},
      ResponseStruct: &armchaos.TargetTypesClientListResponse{},
      Client: &armchaos.TargetTypesClient{},
      ListFunc: (&armchaos.TargetTypesClient{}).NewListPager,
			NewFunc: armchaos.NewTargetTypesClient,
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.Chaos/locations/{locationName}/targetTypes",
		},
		{
      Name: "target",
      Struct: &armchaos.Target{},
      ResponseStruct: &armchaos.TargetsClientListResponse{},
      Client: &armchaos.TargetsClient{},
      ListFunc: (&armchaos.TargetsClient{}).NewListPager,
			NewFunc: armchaos.NewTargetsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{parentProviderNamespace}/{parentResourceType}/{parentResourceName}/providers/Microsoft.Chaos/targets",
		},
	}

	for i := range tables {
		tables[i].Service = "armchaos"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armchaos()...)
}