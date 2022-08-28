package codegen

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"google.golang.org/api/iam/v1"
)

var iamResources = []*Resource{
	{
		SubService: "roles",
		Struct:     &iam.Role{},
	},
	{
		SubService:  "service_accounts",
		Struct:      &iam.ServiceAccount{},
		OutputField: "Accounts",
	},
}

func IamResources() []*Resource {
	var resources []*Resource
	resources = append(resources, iamResources...)

	for _, resource := range resources {
		resource.Service = "iam"
		resource.ListFunction = fmt.Sprintf(`c.Services.Iam.Projects.%s.List("projects/" + c.ProjectId).PageToken(nextPageToken).Do()`, strcase.ToCamel(resource.SubService))
		resource.Template = "resource_list"
		if resource.OutputField == "" {
			resource.OutputField = strcase.ToCamel(resource.SubService)
		}
	}

	return resources
}
