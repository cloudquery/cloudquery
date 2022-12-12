package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/iam/v1"
)



func init() {
	resources := []*Resource{
		{
			SubService: "roles",
			Struct:     &iam.Role{},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "project_id",
					Type:     schema.TypeString,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: "client.ResolveProject",
				},
				{
					Name:    "name",
					Type:    schema.TypeString,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService:  "service_accounts",
			Struct:      &iam.ServiceAccount{},
			OutputField: "Accounts",
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "unique_id",
					Type:     schema.TypeString,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: `schema.PathResolver("UniqueId")`,
				},
			},
			SkipFields:      []string{"ProjectId"},
			NameTransformer: CreateReplaceTransformer(map[string]string{"oauth_2": "oauth2"}),
			Relations:       []string{"ServiceAccountKeys()"},
			SkipMock:        true,
		},
		{
			SubService:  "service_account_keys",
			Struct:      &iam.ServiceAccountKey{},
			ChildTable:  true,
			OutputField: "AccountKeys",
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "service_account_unique_id",
					Type:     schema.TypeString,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: `schema.ParentColumnResolver("unique_id")`,
				},
			},
			SkipFields: []string{"ProjectId", "PrivateKeyData", "PrivateKeyType"},
			SkipMock:   true,
		},
	}

	for _, resource := range resources {
		resource.Service = "iam"
		resource.SkipFetch = true
		resource.MockImports = []string{"google.golang.org/api/iam/v1"}
		resource.Template = "newapi_list"
		resource.MockTemplate = "resource_list_mock"
		if resource.OutputField == "" {
			resource.OutputField = strcase.ToCamel(resource.SubService)
		}
	}

	Resources = append(Resources, resources...)
}
