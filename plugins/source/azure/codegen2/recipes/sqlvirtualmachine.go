// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"

func init() {
	tables := []Table{
		{
			Service:        "armsqlvirtualmachine",
			Name:           "groups",
			Struct:         &armsqlvirtualmachine.Group{},
			ResponseStruct: &armsqlvirtualmachine.GroupsClientListResponse{},
			Client:         &armsqlvirtualmachine.GroupsClient{},
			ListFunc:       (&armsqlvirtualmachine.GroupsClient{}).NewListPager,
			NewFunc:        armsqlvirtualmachine.NewGroupsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_SqlVirtualMachine)`,
		},
		{
			Service:        "armsqlvirtualmachine",
			Name:           "sql_virtual_machines",
			Struct:         &armsqlvirtualmachine.SQLVirtualMachine{},
			ResponseStruct: &armsqlvirtualmachine.SQLVirtualMachinesClientListResponse{},
			Client:         &armsqlvirtualmachine.SQLVirtualMachinesClient{},
			ListFunc:       (&armsqlvirtualmachine.SQLVirtualMachinesClient{}).NewListPager,
			NewFunc:        armsqlvirtualmachine.NewSQLVirtualMachinesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_SqlVirtualMachine)`,
		},
	}
	Tables = append(Tables, tables...)
}
