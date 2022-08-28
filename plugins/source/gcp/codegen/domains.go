package codegen

import (
	"fmt"

	"github.com/iancoleman/strcase"
	domains "google.golang.org/api/domains/v1beta1"
)

var domainsResources = []*Resource{
	{
		SubService: "registrations",
		Struct:     &domains.Registration{},
	},
}

func DomainsResources() []*Resource {
	var resources []*Resource
	resources = append(resources, domainsResources...)

	for _, resource := range resources {
		resource.Service = "domains"
		resource.ListFunction = fmt.Sprintf(
			`c.Services.Domain.Projects.Locations.%s.List("projects/" + c.ProjectId + "/locations/-").PageToken(nextPageToken).Do()`,
			strcase.ToCamel(resource.SubService),
		)
		resource.Template = "resource_list"
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}

	return resources
}
