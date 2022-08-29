package codegen

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	domains "google.golang.org/api/domains/v1beta1"
)

var domainsResources = []*Resource{
	{
		SubService: "registrations",
		Struct:     &domains.Registration{},
		OverrideColumns: []codegen.ColumnDefinition{
			{
				Name:    "name",
				Type:    schema.TypeString,
				Options: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	},
}

func DomainsResources() []*Resource {
	var resources []*Resource
	resources = append(resources, domainsResources...)

	for _, resource := range resources {
		resource.Service = "domains"
		resource.MockImports = []string{"google.golang.org/api/domains/v1beta1"}
		resource.ListFunction = fmt.Sprintf(
			`c.Services.Domains.Projects.Locations.%s.List("projects/" + c.ProjectId + "/locations/-").PageToken(nextPageToken).Do()`,
			strcase.ToCamel(resource.SubService),
		)
		resource.Template = "resource_list"
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}

	return resources
}
