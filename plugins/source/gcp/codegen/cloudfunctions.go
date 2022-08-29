package codegen

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/cloudfunctions/v1"
)

var cloudFunctionsResources = []*Resource{
	{
		SubService: "functions",
		Struct:     &cloudfunctions.CloudFunction{},
		OverrideColumns: []codegen.ColumnDefinition{
			{
				Name:    "name",
				Type:    schema.TypeString,
				Options: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	},
}

func CloudFunctionsResources() []*Resource {
	var resources []*Resource
	resources = append(resources, cloudFunctionsResources...)

	for _, resource := range resources {
		resource.MockImports = []string{"google.golang.org/api/cloudfunctions/v1"}
		resource.Service = "cloudfunctions"
		resource.ListFunction = fmt.Sprintf(`c.Services.Cloudfunctions.Projects.Locations.%s.List("projects/" + c.ProjectId + "/locations/-").PageToken(nextPageToken).Do()`, strcase.ToCamel(resource.SubService))
		resource.Template = "resource_list"
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}
	return resources
}
