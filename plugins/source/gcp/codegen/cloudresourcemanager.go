package codegen

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/cloudresourcemanager/v3"
)

var cloudResourceManagerResources = []*Resource{
	{
		SubService:  "folders",
		Struct:      &cloudresourcemanager.Folder{},
		OutputField: "Folders",
		DefaultColumns: []codegen.ColumnDefinition{
			ProjectIdColumn,
			{
				Name:     "policy",
				Type:     schema.TypeJSON,
				Resolver: "resolveFolderPolicy",
			},
		},
	},
	{
		SubService: "projects",
		Struct:     &cloudresourcemanager.Project{},
		DefaultColumns: []codegen.ColumnDefinition{
			{
				Name:     "policy",
				Type:     schema.TypeJSON,
				Resolver: "resolveProjectPolicy",
			},
		},
		ListFunction: `c.Services.Cloudresourcemanager.Projects.Get("projects/" + c.ProjectId).Do()`,
	},
}

func CloudResourceManagerResources() []*Resource {
	var resources []*Resource
	resources = append(resources, cloudResourceManagerResources...)

	for _, resource := range resources {
		resource.SkipMock = true
		resource.Service = "cloudresourcemanager"
		resource.Template = "resource_get"
		if resource.ListFunction == "" {
			resource.ListFunction = fmt.Sprintf(`c.Services.Cloudresourcemanager.%s.List().Do()`, strcase.ToCamel(resource.SubService))
		}
	}

	return resources
}
