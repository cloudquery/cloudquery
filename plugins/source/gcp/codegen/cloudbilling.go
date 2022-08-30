package codegen

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/cloudbilling/v1"
)

var cloudbillingResources = []*Resource{
	{
		SubService: "billing_accounts",
		Struct:     &cloudbilling.BillingAccount{},
		OverrideColumns: []codegen.ColumnDefinition{
			{
				Name:    "name",
				Type:    schema.TypeString,
				Options: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	},
	{
		SubService: "services",
		Struct:     &cloudbilling.Service{},
		OverrideColumns: []codegen.ColumnDefinition{
			{
				Name:    "name",
				Type:    schema.TypeString,
				Options: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	},
}

func CloudBillingResources() []*Resource {
	var resources []*Resource
	resources = append(resources, cloudbillingResources...)

	for _, resource := range resources {
		resource.Service = "cloudbilling"
		resource.Template = "resource_list"
		resource.MockImports = []string{"google.golang.org/api/cloudbilling/v1"}
		resource.ListFunction = fmt.Sprintf(`c.Services.Cloudbilling.%s.List().PageToken(nextPageToken).Do()`, strcase.ToCamel(resource.SubService))
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}

	return resources
}
