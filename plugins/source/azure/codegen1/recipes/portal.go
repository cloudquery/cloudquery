// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"

func Armportal() []*Table {
	tables := []*Table{
		{
			NewFunc: armportal.NewListTenantConfigurationViolationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal",
			URL:     "/providers/Microsoft.Portal/listTenantConfigurationViolations",
		},
		{
			NewFunc: armportal.NewTenantConfigurationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal",
			URL:     "/providers/Microsoft.Portal/tenantConfigurations",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armportal())
}
