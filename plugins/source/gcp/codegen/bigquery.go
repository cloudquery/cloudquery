package codegen

import (
	"cloud.google.com/go/bigquery"
)

var bigqueryResources = []*Resource{
	{
		SubService:          "tables",
		Struct:              &bigquery.TableMetadata{},
		NewFunction:         bigquery.NewClient,
		SkipFetch:           true,
		PreResourceResolver: "tableGet",
		// GetFunction:  "c.Services.Bigquery.Tables.Get(c.ProjectId, r.Item.(*bigquery.TableListTables).TableReference.DatasetId, r.Item.(*bigquery.TableListTables).TableReference.TableId).Do()",
		Imports:    []string{"google.golang.org/api/bigquery/v2"},
		Multiplex:  &emptyString,
		ChildTable: true,
		SkipMock:   true,
	},
	{
		SubService:          "datasets",
		Struct:              &bigquery.DatasetMetadata{},
		NewFunction:         bigquery.NewClient,
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
		resource.MockImports = []string{"cloud.google.com/go/bigquery"}
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
	}

	return resources
}
