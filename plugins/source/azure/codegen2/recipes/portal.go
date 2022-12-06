// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"

func Armportal() []Table {
	tables := []Table{
		{
      Name: "configuration",
      Struct: &armportal.Configuration{},
      ResponseStruct: &armportal.TenantConfigurationsClientListResponse{},
      Client: &armportal.TenantConfigurationsClient{},
      ListFunc: (&armportal.TenantConfigurationsClient{}).NewListPager,
			NewFunc: armportal.NewTenantConfigurationsClient,
			URL: "/providers/Microsoft.Portal/tenantConfigurations",
		},
		{
      Name: "violation",
      Struct: &armportal.Violation{},
      ResponseStruct: &armportal.ListTenantConfigurationViolationsClientListResponse{},
      Client: &armportal.ListTenantConfigurationViolationsClient{},
      ListFunc: (&armportal.ListTenantConfigurationViolationsClient{}).NewListPager,
			NewFunc: armportal.NewListTenantConfigurationViolationsClient,
			URL: "/providers/Microsoft.Portal/listTenantConfigurationViolations",
		},
	}

	for i := range tables {
		tables[i].Service = "armportal"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
  Tables = append(Tables, Armportal()...)
}