// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"

func Armcontainerregistry() []Table {
	tables := []Table{
		{
			Name:           "registry",
			Struct:         &armcontainerregistry.Registry{},
			ResponseStruct: &armcontainerregistry.RegistriesClientListResponse{},
			Client:         &armcontainerregistry.RegistriesClient{},
			ListFunc:       (&armcontainerregistry.RegistriesClient{}).NewListPager,
			NewFunc:        armcontainerregistry.NewRegistriesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerRegistry/registries",
		},
	}

	for i := range tables {
		tables[i].Service = "armcontainerregistry"
		tables[i].Template = "list"
	}
	return tables
}

func init() {
	Tables = append(Tables, Armcontainerregistry()...)
}
