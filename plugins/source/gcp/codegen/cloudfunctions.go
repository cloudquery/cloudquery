package codegen

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"google.golang.org/api/cloudfunctions/v1"
)

var cloudFunctionsResources = []*Resource{
	{
		SubService: "functions",
		Struct:     &cloudfunctions.CloudFunction{},
	},
}

func CloudFunctionsResources() []*Resource {
	var resources []*Resource
	resources = append(resources, cloudFunctionsResources...)

	for _, resource := range resources {
		resource.Service = "cloudfunctions"
		resource.ListFunction = fmt.Sprintf(`c.Services.CloudFunctions.Projects.Locations.%s.List("projects/" + c.ProjectId + "/locations/-").PageToken(nextPageToken).Do()`, strcase.ToCamel(resource.SubService))
		resource.Template = "resource_list"
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}
	return resources
}
