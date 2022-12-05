package recipes

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Users() []*Resource {
	resources := []*Resource{
		{
			SubService: "users",
			Multiplex:  "client.AccountMultiplex",
			Struct:     new(datadogV2.User),
			SkipFields: []string{"Id"},
			ExtraColumns: append(defaultAccountColumnsPK, codegen.ColumnDefinition{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("Id")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
			),
			Relations: []string{"UserPermissions()", "UserOrganizations()"},
		},
		{
			SubService:   "user_permissions",
			Struct:       new(datadogV2.Permission),
			ExtraColumns: defaultAccountColumns,
		},
		{
			SubService:   "user_organizations",
			Struct:       new(datadogV2.User),
			ExtraColumns: defaultAccountColumns,
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "users"
	}
	return resources
}
