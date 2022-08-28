package codegen

import (
	"fmt"
	"reflect"

	"github.com/cloudquery/plugin-sdk/codegen"
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
		resource.Service = "bigquery"
		resource.DefaultColumns = []codegen.ColumnDefinition{ProjectIdColumn}
		resource.StructName = reflect.TypeOf(resource.Struct).Elem().Name()
		resource.ListFunction = fmt.Sprintf("c.Services.BigQuery.%s.List(c.ProjectId).PageToken(nextPageToken).Do()", strcase.ToCamel(resource.SubService))
		if resource.Template == "" {
			resource.Template = "resource_list"
		}
		if resource.SkipFields == nil {
			resource.SkipFields = []string{"ServerResponse", "NullFields", "ForceSendFields"}
		}
		resource.MockImports = []string{"google.golang.org/api/bigquery/v2"}
		if resource.MockListStruct == "" {
			resource.MockListStruct = strcase.ToCamel(resource.StructName)
		}
	}

	return resources
}
