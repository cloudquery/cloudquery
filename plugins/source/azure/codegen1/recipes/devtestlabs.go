// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"

func Armdevtestlabs() []*Table {
	tables := []*Table{
		{
			NewFunc: armdevtestlabs.NewProviderOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs",
			URL:     "/providers/Microsoft.DevTestLab/operations",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdevtestlabs())
}
