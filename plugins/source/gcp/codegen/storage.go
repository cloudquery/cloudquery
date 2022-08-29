package codegen

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"google.golang.org/api/storage/v1"
)

var storageResources = []*Resource{
	{
		SubService: "buckets",
		Struct:     &storage.Bucket{},
		OverrideColumns: []codegen.ColumnDefinition{
			{
				Name:    "self_link",
				Type:    schema.TypeString,
				Options: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	},
}

func StorageResources() []*Resource {
	var resources []*Resource
	resources = append(resources, storageResources...)

	for _, resource := range resources {
		resource.Service = "storage"
		resource.MockImports = []string{"google.golang.org/api/storage/v1"}
		resource.Template = "resource_list"
		resource.ListFunction = fmt.Sprintf("c.Services.Storage.%s.List(c.ProjectId).PageToken(nextPageToken).Do()", strcase.ToCamel(resource.SubService))
	}

	return resources
}
