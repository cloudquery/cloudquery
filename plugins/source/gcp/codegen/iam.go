package codegen

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/iam/v1"
)

var iamResources = []*Resource{
	{
		SubService:   "roles",
		Struct:       &iam.Role{},
		NewFunction:  iam.NewProjectsRolesService,
		ListFunction: (&iam.ProjectsRolesService{}).List,
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
		SubService:   "service_accounts",
		Struct:       &iam.ServiceAccount{},
		NewFunction:  iam.NewProjectsServiceAccountsService,
		ListFunction: (&iam.ProjectsServiceAccountsService{}).List,
		OutputField:  "Accounts",
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
	},
}

func IamResources() []*Resource {
	var resources []*Resource
	resources = append(resources, iamResources...)

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

	return resources
}
