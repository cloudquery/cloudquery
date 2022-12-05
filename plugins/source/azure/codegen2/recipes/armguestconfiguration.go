// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"

func Armguestconfiguration() []Table {
	tables := []Table{
		{
      Name: "assignment",
      Struct: &armguestconfiguration.Assignment{},
      ResponseStruct: &armguestconfiguration.HCRPAssignmentsClientListResponse{},
      Client: &armguestconfiguration.HCRPAssignmentsClient{},
      ListFunc: (&armguestconfiguration.HCRPAssignmentsClient{}).NewListPager,
			NewFunc: armguestconfiguration.NewHCRPAssignmentsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments",
		},
		{
      Name: "assignment",
      Struct: &armguestconfiguration.Assignment{},
      ResponseStruct: &armguestconfiguration.AssignmentsClientListResponse{},
      Client: &armguestconfiguration.AssignmentsClient{},
      ListFunc: (&armguestconfiguration.AssignmentsClient{}).NewListPager,
			NewFunc: armguestconfiguration.NewAssignmentsClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments",
		},
		{
      Name: "assignment",
      Struct: &armguestconfiguration.Assignment{},
      ResponseStruct: &armguestconfiguration.AssignmentsVMSSClientListResponse{},
      Client: &armguestconfiguration.AssignmentsVMSSClient{},
      ListFunc: (&armguestconfiguration.AssignmentsVMSSClient{}).NewListPager,
			NewFunc: armguestconfiguration.NewAssignmentsVMSSClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmssName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments",
		},
		{
      Name: "assignment_report",
      Struct: &armguestconfiguration.AssignmentReport{},
      ResponseStruct: &armguestconfiguration.AssignmentReportsVMSSClientListResponse{},
      Client: &armguestconfiguration.AssignmentReportsVMSSClient{},
      ListFunc: (&armguestconfiguration.AssignmentReportsVMSSClient{}).NewListPager,
			NewFunc: armguestconfiguration.NewAssignmentReportsVMSSClient,
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmssName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{name}/reports",
		},
	}

	for i := range tables {
		tables[i].Service = "armguestconfiguration"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armguestconfiguration()...)
}