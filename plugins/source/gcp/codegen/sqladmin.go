package codegen

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

var sqlResources = []*Resource{
	{
		SubService: "instances",
		Struct:     &sqladmin.DatabaseInstance{},
		OverrideColumns: []codegen.ColumnDefinition{
			{
				Name:     "self_link",
				Type:     schema.TypeString,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				Resolver: `schema.PathResolver("SelfLink")`,
			},
		},
	},
}

func SqlResources() []*Resource {
	var resources []*Resource
	resources = append(resources, sqlResources...)

	for _, resource := range resources {
		resource.Service = "sqladmin"
		resource.MockImports = []string{"google.golang.org/api/sqladmin/v1beta4"}
		resource.Template = "resource_list"
		resource.ListFunction = fmt.Sprintf(`c.Services.Sqladmin.%s.List(c.ProjectId).PageToken(nextPageToken).Do()`, strcase.ToCamel(resource.SubService))
	}

	return resources
}
