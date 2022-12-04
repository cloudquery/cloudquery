// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata"

func Armazuredata() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armazuredata.NewSQLServersClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata",
		},
		{
			NewFunc: armazuredata.NewOperationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata",
		},
		{
			NewFunc: armazuredata.NewSQLServerRegistrationsClient,
			PkgPath: "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata",
		},
	}
	return resources
}

func init() {
	Resources = append(Resources, Armazuredata())
}