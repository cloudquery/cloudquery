package codegen

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/cloudresourcemanager/v3"
)

var cloudResourceManagerResources = []*Resource{
	{
		SubService:  "folders",
		Struct:      &cloudresourcemanager.Folder{},
		OutputField: "Folders",
	},
	{
		SubService:     "projects",
		Struct:         &cloudresourcemanager.Project{},
		DefaultColumns: make([]codegen.ColumnDefinition, 0),
		OutputField:    "Projects",
	},
}

func CloudResourceManagerResources() []*Resource {
	var resources []*Resource
	resources = append(resources, cloudResourceManagerResources...)

	for _, resource := range resources {
		resource.Service = "cloudresourcemanager"
		resource.Template = "resource_get"
		resource.ListFunction = fmt.Sprintf(`c.Services.Cloudresourcemanager.%s.List().Do()`, strcase.ToCamel(resource.SubService))
	}

	return resources
}
