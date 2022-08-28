package codegen

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"google.golang.org/api/cloudbilling/v1"
)

var cloudbillingResources = []*Resource{
	{
		SubService: "billing_accounts",
		Struct:     &cloudbilling.BillingAccount{},
	},
}

func CloudBillingResources() []*Resource {
	var resources []*Resource
	resources = append(resources, cloudbillingResources...)

	for _, resource := range resources {
		resource.Service = "cloudbilling"
		resource.Template = "resource_list"
		resource.ListFunction = fmt.Sprintf(`c.Services.CloudBilling.%s.List().PageToken(nextPageToken).Do()`, strcase.ToCamel(resource.SubService))
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}

	return resources
}
