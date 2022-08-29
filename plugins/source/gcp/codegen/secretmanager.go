package codegen

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/secretmanager/v1"
)

var secretmanagerResources = []*Resource{
	{
		SubService: "secrets",
		Struct:     &secretmanager.Secret{},
		OverrideColumns: []codegen.ColumnDefinition{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				Resolver: `schema.PathResolver("Name")`,
			},
		},
	},
}

func SecretManagerResources() []*Resource {
	var resources []*Resource
	resources = append(resources, secretmanagerResources...)

	for _, resource := range resources {
		resource.Service = "secretmanager"
		resource.Template = "resource_list"
		resource.ListFunction = fmt.Sprintf(
			`c.Services.Secretmanager.Projects.%s.List("projects/" + c.ProjectId).PageToken(nextPageToken).Do()`,
			strcase.ToCamel(resource.SubService),
		)
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}

	return resources
}
