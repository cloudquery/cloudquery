package recipes

import (
	"cloud.google.com/go/iam"
	"cloud.google.com/go/storage"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func init() {
	resources := []*Resource{
		{
			SubService:      "buckets",
			Struct:          &storage.BucketAttrs{},
			SkipFetch:       true,
			SkipMock:        true,
			NameTransformer: CreateReplaceTransformer(map[string]string{"c_o_r_s": "cors", "r_p_o": "rpo"}),
			PrimaryKeys:     []string{"name"},
			Relations:       []string{"BucketPolicies()"},
			Description:     "https://cloud.google.com/storage/docs/json_api/v1/buckets#resource",
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
			Description: "https://cloud.google.com/iam/docs/reference/rest/v1/Policy",
		},
	}

	for _, resource := range resources {
		resource.Service = "storage"
		resource.MockImports = []string{"cloud.google.com/go/storage"}
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_rest_mock"
	}

	Resources = append(Resources, resources...)
}
