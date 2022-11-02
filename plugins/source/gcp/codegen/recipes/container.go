package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/container/v1"
)

var containerResources = []*Resource{
	{
		SubService: "clusters",
		Struct:     new(container.Cluster),
		SkipFetch:  true,
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:    "self_link",
				Type:    schema.TypeString,
				Options: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
		NameTransformer: CreateReplaceTransformer(map[string]string{"ipv_4": "ipv4"}),
	},
}

func ContainerResources() []*Resource {
	var resources []*Resource
	resources = append(resources, containerResources...)

	for _, resource := range resources {
		resource.Service = "container"
		resource.Template = "newapi_list"
		resource.MockTemplate = "resource_list_mock"
		if resource.OutputField == "" {
			resource.OutputField = strcase.ToCamel(resource.SubService)
		}
	}

	return resources
}
