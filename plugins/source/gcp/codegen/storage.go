package codegen

import (
	"cloud.google.com/go/storage"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var storageResources = []*Resource{
	{
		SubService:     "buckets",
		Struct:         &storage.BucketAttrs{},
		NewFunction:    storage.NewClient,
		ResponseStruct: &storage.BucketAttrs{},
		SkipFetch:      true,
		SkipMock:       true,
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
		resource.MockImports = []string{"cloud.google.com/go/storage"}
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_rest_mock"
	}

	return resources
}
