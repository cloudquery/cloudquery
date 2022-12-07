// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"

func Armsqlvirtualmachine() []Table {
	tables := []Table{
		{
			Name:           "sql_virtual_machines",
			Struct:         &armsqlvirtualmachine.SQLVirtualMachine{},
			ResponseStruct: &armsqlvirtualmachine.SQLVirtualMachinesClientListResponse{},
			Client:         &armsqlvirtualmachine.SQLVirtualMachinesClient{},
			ListFunc:       (&armsqlvirtualmachine.SQLVirtualMachinesClient{}).NewListPager,
			NewFunc:        armsqlvirtualmachine.NewSQLVirtualMachinesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.SqlVirtualMachine")`,
		},
		{
			Name:           "groups",
			Struct:         &armsqlvirtualmachine.Group{},
			ResponseStruct: &armsqlvirtualmachine.GroupsClientListResponse{},
			Client:         &armsqlvirtualmachine.GroupsClient{},
			ListFunc:       (&armsqlvirtualmachine.GroupsClient{}).NewListPager,
			NewFunc:        armsqlvirtualmachine.NewGroupsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace("Microsoft.SqlVirtualMachine")`,
		},
	}

	for i := range tables {
		tables[i].Service = "armsqlvirtualmachine"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armsqlvirtualmachine()...)
}
