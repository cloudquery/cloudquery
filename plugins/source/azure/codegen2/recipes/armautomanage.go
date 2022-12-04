// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage"

func Armautomanage() []Table {
	tables := []Table{
		{
      Name: "configuration_profile_assignment",
      Struct: &armautomanage.ConfigurationProfileAssignment{},
      ResponseStruct: &armautomanage.ConfigurationProfileAssignmentsClientListResponse{},
      Client: &armautomanage.ConfigurationProfileAssignmentsClient{},
      ListFunc: (&armautomanage.ConfigurationProfileAssignmentsClient{}).NewListPager,
			NewFunc: armautomanage.NewConfigurationProfileAssignmentsClient,
		},
		{
      Name: "operation",
      Struct: &armautomanage.Operation{},
      ResponseStruct: &armautomanage.OperationsClientListResponse{},
      Client: &armautomanage.OperationsClient{},
      ListFunc: (&armautomanage.OperationsClient{}).NewListPager,
			NewFunc: armautomanage.NewOperationsClient,
		},
	}

	for i := range tables {
		tables[i].Service = "armautomanage"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armautomanage()...)
}