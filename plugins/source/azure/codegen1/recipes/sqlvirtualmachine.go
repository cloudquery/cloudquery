// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine"

func Armsqlvirtualmachine() []*Table {
	tables := []*Table{
		{
			NewFunc: armsqlvirtualmachine.NewGroupsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachineGroups",
		},
		{
			NewFunc: armsqlvirtualmachine.NewSQLVirtualMachinesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sqlvirtualmachine/armsqlvirtualmachine",
			URL:     "/subscriptions/{subscriptionId}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armsqlvirtualmachine())
}
