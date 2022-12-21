package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/iam/v1"
	iamv2Beta "google.golang.org/api/iam/v2beta"
)

func init() {
	resources := []*Resource{
		{
			SubService:  "roles",
			Struct:      &iam.Role{},
			PrimaryKeys: []string{ProjectIdColumn.Name, "name"},
			Description: "https://cloud.google.com/iam/docs/reference/rest/v1/roles#Role",
			MockImports: []string{"google.golang.org/api/iam/v1"},
		},
		{
			SubService:      "service_accounts",
			Struct:          &iam.ServiceAccount{},
			OutputField:     "Accounts",
			PrimaryKeys:     []string{"unique_id"},
			SkipFields:      []string{"ProjectId"},
			NameTransformer: CreateReplaceTransformer(map[string]string{"oauth_2": "oauth2"}),
			Relations:       []string{"ServiceAccountKeys()"},
			SkipMock:        true,
			Description:     "https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts#ServiceAccount",
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
			SkipFields:  []string{"ProjectId", "PrivateKeyData", "PrivateKeyType"},
			SkipMock:    true,
			Description: "https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKey",
		},
		{
			SubService:  "deny_policies",
			SkipMock:    true,
			Struct:      &iamv2Beta.GoogleIamV2betaPolicy{},
			Description: "https://cloud.google.com/iam/docs/reference/rest/v2beta/policies#Policy",
		},
	}

	for _, resource := range resources {
		resource.Service = "iam"
		resource.SkipFetch = true
		resource.Template = "newapi_list"
		resource.MockTemplate = "resource_list_mock"
		if resource.OutputField == "" {
			resource.OutputField = strcase.ToCamel(resource.SubService)
		}
	}

	Resources = append(Resources, resources...)
}
