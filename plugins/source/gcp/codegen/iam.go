package codegen

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/iam/v1"
)

var iamResources = []*Resource{
	{
		SubService: "roles",
		Struct:     &iam.Role{},
		OverrideColumns: []codegen.ColumnDefinition{
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
		SkipFields: []string{"ServerResponse", "NullFields", "ForceSendFields"},
	},
	{
		SubService:  "service_accounts",
		Struct:      &iam.ServiceAccount{},
		OutputField: "Accounts",
		OverrideColumns: []codegen.ColumnDefinition{
			{
				Name:     "unique_id",
				Type:     schema.TypeString,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				Resolver: `schema.PathResolver("UniqueId")`,
			},
		},
		SkipFields: []string{"ProjectId", "NullFields", "ForceSendFields"},
	},
}

func IamResources() []*Resource {
	var resources []*Resource
	resources = append(resources, iamResources...)

	for _, resource := range resources {
		resource.Service = "iam"
		resource.MockImports = []string{"google.golang.org/api/iam/v1"}
		resource.ListFunction = fmt.Sprintf(`c.Services.Iam.Projects.%s.List("projects/" + c.ProjectId).PageToken(nextPageToken).Do()`, strcase.ToCamel(resource.SubService))
		resource.Template = "resource_list"
		if resource.OutputField == "" {
			resource.OutputField = strcase.ToCamel(resource.SubService)
		}
	}

	return resources
}
