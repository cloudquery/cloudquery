package codegen

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"google.golang.org/api/cloudresourcemanager/v3"
)

var cloudResourceManagerResources = []*Resource{
	{
		SubService: "folders",
		Struct:     &cloudresourcemanager.Folder{},
	},
}

func CloudResourceManagerResources() []*Resource {
	var resources []*Resource
	resources = append(resources, cloudResourceManagerResources...)

	for _, resource := range resources {
		resource.Service = "cloudresourcemanager"
		resource.Template = "resource_list"
		resource.ListFunction = fmt.Sprintf(`c.Services.ResourceManager.%s.List().Do()`, strcase.ToCamel(resource.SubService))
	}

	return resources
}
