package recipies

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Users() []*Resource {
	return []*Resource{
		{
			Service:    "users",
			SubService: "users",
			Multiplex:  "client.AccountMultiplex",
			Struct:     new(datadogV2.User),
			SkipFields: []string{"Id"},
			ExtraColumns: append(defaultAccountColumns, codegen.ColumnDefinition{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: `schema.PathResolver("Id")`,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
			},
			),
			Relations: []string{"UserPermissions()", "UserOrganizations()"},
		},
		{
			Service:      "users",
			SubService:   "user_permissions",
			Struct:       new(datadogV2.Permission),
			ExtraColumns: defaultAccountColumns,
		},
		{
			Service:      "users",
			SubService:   "user_organizations",
			Struct:       new(datadogV2.User),
			ExtraColumns: defaultAccountColumns,
		},
	}
}
