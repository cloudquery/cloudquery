package codegen

import (
	"github.com/iancoleman/strcase"
	"google.golang.org/api/bigquery/v2"
)

var bigqueryResources = []*Resource{
	{
		SubService:   "tables",
		Struct:       &bigquery.Table{},
		ListFunction: "c.Services.Bigquery.Tables.List(c.ProjectId, r.Parent.Item.(*bigquery.DatasetListDatasets).DatasetReference.DatasetId).PageToken(nextPageToken).Do()",
		GetFunction:  "c.Services.Bigquery.Tables.Get(c.ProjectId, r.Item.(*bigquery.TableListTables).TableReference.DatasetId, r.Item.(*bigquery.TableListTables).TableReference.TableId).Do()",
		Imports:      []string{"google.golang.org/api/bigquery/v2"},
		Multiplex:    &emptyString,
		ChildTable:   true,
		SkipMock:     true,
	},
	{
		SubService:   "datasets",
		Struct:       &bigquery.Dataset{},
		ListFunction: "c.Services.Bigquery.Datasets.List(c.ProjectId).PageToken(nextPageToken).Do()",
		GetFunction:  "c.Services.Bigquery.Datasets.Get(c.ProjectId, r.Item.(*bigquery.DatasetListDatasets).DatasetReference.DatasetId).Do()",
		Imports:      []string{"google.golang.org/api/bigquery/v2"},
		Relations:    []string{"Tables()"},
		SkipMock:     true,
	},
}

func BigqueryResources() []*Resource {
	var resources []*Resource
	resources = append(resources, bigqueryResources...)

	for _, resource := range resources {
		resource.MockImports = []string{"google.golang.org/api/bigquery/v2"}
		resource.Service = "bigquery"
		resource.Template = "resource_list"
		resource.OutputField = strcase.ToCamel(resource.SubService)
	}

	return resources
}
