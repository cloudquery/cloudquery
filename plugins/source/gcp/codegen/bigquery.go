package codegen

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"google.golang.org/api/bigquery/v2"
)

var bigqueryResources = []*Resource{
	{
		SubService: "datasets",
		Struct:     &bigquery.DatasetListDatasets{},
	},
}

func BigqueryResources() []*Resource {
	var resources []*Resource
	resources = append(resources, bigqueryResources...)

	for _, resource := range resources {
		resource.MockImports = []string{"google.golang.org/api/bigquery/v2"}
		resource.Service = "bigquery"
		resource.ListFunction = fmt.Sprintf("c.Services.Bigquery.%s.List(c.ProjectId).PageToken(nextPageToken).Do()", strcase.ToCamel(resource.SubService))
		resource.Template = "resource_list"
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}

	return resources
}
