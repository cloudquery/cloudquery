package recipes

import (
	bigquery "google.golang.org/api/bigquery/v2"
)


func init() {
	resources := []*Resource{
		{
			SubService:          "tables",
			Struct:              &bigquery.Table{},
			SkipFetch:           true,
			PreResourceResolver: "tableGet",
			Multiplex:           &emptyString,
			ChildTable:          true,
			SkipMock:            true,
		},
		{
			SubService:          "datasets",
			Struct:              &bigquery.Dataset{},
			SkipFetch:           true,
			PreResourceResolver: "datasetGet",
			Relations:           []string{"Tables()"},
			SkipMock:            true,
		},
	}

	for _, resource := range resources {
		resource.Service = "bigquery"
		resource.Template = "newapi_list"
	}

	Resources = append(Resources, resources...)
}
