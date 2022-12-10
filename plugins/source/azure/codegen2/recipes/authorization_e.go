package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization"

func ArmauthorizationE() []Table {
	tables := []Table{
		{
			Name:           "role_definitions",
			Struct:         &armauthorization.RoleDefinition{},
			ResponseStruct: &armauthorization.RoleDefinitionsClientListResponse{},
			Client:         &armauthorization.RoleDefinitionsClient{},
			ListFunc:       (&armauthorization.RoleDefinitionsClient{}).NewListPager,
			NewFunc:        armauthorization.NewRoleDefinitionsClient,
			URL:            "/{scope}/providers/Microsoft.Authorization/roleDefinitions",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Authorization)`,
			SkipFetch:      true,
		},
	}
	for i := range tables {
		tables[i].Service = "armauthorization"
	}
	return tables
}

func init() {
	Tables = append(Tables, ArmauthorizationE()...)
}
