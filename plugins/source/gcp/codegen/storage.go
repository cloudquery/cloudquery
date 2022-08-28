package codegen

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"google.golang.org/api/storage/v1"
)

var storageResources = []*Resource{
	{
		SubService: "buckets",
		Struct:     &storage.Bucket{},
	},
}

func StorageResources() []*Resource {
	var resources []*Resource
	resources = append(resources, storageResources...)

	for _, resource := range resources {
		resource.Service = "storage"
		resource.Template = "resource_list"
		resource.ListFunction = fmt.Sprintf("c.Services.Storage.%s.List(c.ProjectId).PageToken(nextPageToken).Do()", strcase.ToCamel(resource.SubService))
	}

	return resources
}
