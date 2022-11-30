package recipes

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Roles() []*Resource {
	resources := []*Resource{
		{
			SubService: "permissions",
			Multiplex:  "client.AccountMultiplex",
			Struct:     new(datadogV2.Permission),
			SkipFields: []string{"Id"},
			ExtraColumns: append(defaultAccountColumnsPK, codegen.ColumnDefinition{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("Id")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
			),
		},
		{
			SubService: "roles",
			Multiplex:  "client.AccountMultiplex",
			Struct:     new(datadogV2.Role),
			SkipFields: []string{"Id"},
			ExtraColumns: append(defaultAccountColumnsPK, codegen.ColumnDefinition{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("Id")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
			),
			Relations: []string{"RolePermissions()", "RoleUsers()"},
		},
		{
			SubService:   "role_permissions",
			Struct:       new(datadogV2.Permission),
			ExtraColumns: defaultAccountColumns,
		},
		{
			SubService:   "role_users",
			Struct:       new(datadogV2.User),
			ExtraColumns: defaultAccountColumns,
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "roles"
	}
	return resources
}
