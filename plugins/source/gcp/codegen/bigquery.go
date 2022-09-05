package codegen

import (
	"cloud.google.com/go/bigquery"
)

var bigqueryResources = []*Resource{
	{
		SubService:          "tables",
		Struct:              &bigquery.TableMetadata{},
		SkipFetch:           true,
		PreResourceResolver: "tableGet",
		Multiplex:           &emptyString,
		ChildTable:          true,
		SkipMock:            true,
	},
	{
		SubService:          "datasets",
		Struct:              &bigquery.DatasetMetadata{},
		SkipFetch:           true,
		PreResourceResolver: "datasetGet",
		Relations:           []string{"Tables()"},
		SkipMock:            true,
	},
}

func BigqueryResources() []*Resource {
	var resources []*Resource
	resources = append(resources, bigqueryResources...)

	for _, resource := range resources {
		resource.Service = "bigquery"
		resource.Template = "newapi_list"
	}

	return resources
}
