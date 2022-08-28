package codegen

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"google.golang.org/api/secretmanager/v1"
)

var secretmanagerResources = []*Resource{
	{
		SubService: "secrets",
		Struct:     &secretmanager.Secret{},
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
