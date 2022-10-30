package recipes

import (
	"cloud.google.com/go/iam"
	"cloud.google.com/go/storage"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var storageResources = []*Resource{
	{
		SubService:      "buckets",
		Struct:          &storage.BucketAttrs{},
		SkipFetch:       true,
		SkipMock:        true,
		NameTransformer: CreateReplaceTransformer(map[string]string{"c_o_r_s": "cors", "r_p_o": "rpo"}),
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				Resolver: `schema.PathResolver("Name")`,
			},
		},
		Relations: []string{"BucketPolicies()"},
	},
	{
		SubService: "bucket_policies",
		Struct:     &iam.Policy3{},
		SkipFetch:  true,
		SkipMock:   true,
		ChildTable: true,
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:     "bucket_name",
				Type:     schema.TypeString,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				Resolver: `schema.ParentColumnResolver("name")`,
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
