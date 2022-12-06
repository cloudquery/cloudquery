// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices"

func Armdomainservices() []*Table {
	tables := []*Table{
		{
			NewFunc: armdomainservices.NewOuContainerOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices",
			URL: "/providers/Microsoft.Aad/operations",
		},
		{
			NewFunc: armdomainservices.NewDomainServiceOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices",
			URL: "/providers/Microsoft.AAD/operations",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdomainservices())
}