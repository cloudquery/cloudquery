package codegen

import (
	"cloud.google.com/go/storage"
)

var storageResources = []*Resource{
	{
		SubService:     "buckets",
		Struct:         &storage.BucketAttrs{},
		NewFunction:    storage.NewClient,
		ResponseStruct: &storage.BucketAttrs{},
		SkipFetch:      true,
		SkipMock:       true,
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
