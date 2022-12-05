// Code generated by codegen; DO NOT EDIT.
package packages

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork"

func Armmanagednetwork() []*Table {
	tables := []*Table{
		{
			NewFunc: armmanagednetwork.NewScopeAssignmentsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managednetwork/armmanagednetwork",
			URL: "/{scope}/providers/Microsoft.ManagedNetwork/scopeAssignments",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armmanagednetwork())
}