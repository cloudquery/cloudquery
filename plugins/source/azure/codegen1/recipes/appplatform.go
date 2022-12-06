// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"

func Armappplatform() []*Table {
	tables := []*Table{
		{
			NewFunc: armappplatform.NewSKUsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "/subscriptions/{subscriptionId}/providers/Microsoft.AppPlatform/skus",
		},
		{
			NewFunc: armappplatform.NewServicesClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armappplatform())
}