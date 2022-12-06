// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"

func Armquota() []*Table {
	tables := []*Table{
		{
			NewFunc: armquota.NewOperationClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota",
			URL: "/providers/Microsoft.Quota/operations",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armquota())
}